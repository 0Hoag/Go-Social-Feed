package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hoag/go-social-feed/config"
	appMongo "github.com/hoag/go-social-feed/internal/appconfig/mongo"
	"github.com/hoag/go-social-feed/internal/crawler"
	"github.com/hoag/go-social-feed/internal/crawler/sites"
	prod "github.com/hoag/go-social-feed/internal/delivery/rabbitmq/producer"
	"github.com/hoag/go-social-feed/internal/models"
	"github.com/hoag/go-social-feed/internal/post"
	postMongo "github.com/hoag/go-social-feed/internal/post/repository/mongo"
	postUC "github.com/hoag/go-social-feed/internal/post/usecase"
	"github.com/hoag/go-social-feed/internal/processor"
	userMongo "github.com/hoag/go-social-feed/internal/users/repository/mongo"
	userUC "github.com/hoag/go-social-feed/internal/users/usecase"
	pkgCrt "github.com/hoag/go-social-feed/pkg/encrypter"
	pkgLog "github.com/hoag/go-social-feed/pkg/log"
	"github.com/hoag/go-social-feed/pkg/rabbitmq"
	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	// 0. Load .env
	_ = godotenv.Load()

	// 1. Load Config
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	// 2. Logger
	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Logger.Level,
		Mode:     cfg.Logger.Mode,
		Encoding: cfg.Logger.Encoding,
	})

	// 3. Database
	crp := pkgCrt.NewEncrypter(cfg.Encrypter.Key)
	client, err := appMongo.Connect(cfg.Mongo, crp)
	if err != nil {
		l.Fatalf(context.Background(), "MongoDB Connect: %v", err)
	}
	defer client.Disconnect(context.Background())

	db := client.Database(cfg.Mongo.Database)

	// 4. Dependencies
	// RabbitMQ
	amqpConn, err := rabbitmq.Dial(cfg.RabbitConfig.URL, true)
	if err != nil {
		l.Warnf(context.Background(), "RabbitMQ not connected, running without queue...")
		amqpConn = rabbitmq.Connection{}
	}
	defer amqpConn.Close()

	// Producer
	producer := prod.New(l, amqpConn)
	// Optionally run producer loop if needed, though for just posting it might not be strictly required immediately
	// but good practice. `producer.Run()` usually starts a listener for retries or something.
	// Looking at `internal/delivery/rabbitmq/producer/new.go`, `Run` is part of interface.
	// In handler.go it calls `postProd.Run()`.
	if err := producer.Run(); err != nil {
		l.Errorf(context.Background(), "Producer Run failed: %v", err)
	}
	defer producer.Close()

	// Repositories
	userRepo := userMongo.New(l, db)
	postRepo := postMongo.New(l, db)

	// Usecases
	uUC := userUC.New(l, userRepo)
	pUC := postUC.New(l, producer, uUC, postRepo)

	// Crawler & Processor
	crawlMgr := crawler.NewManager(l)
	crawlMgr.Register(sites.NewCoindeskCrawler())
	crawlMgr.Register(sites.NewCoinTelegraphCrawler())

	// Init Gemini Processor
	// We need a context for initialization, but main only has it inside job?
	// Actually NewGeminiProcessor needs context for creating client once.
	proc, err := processor.NewGeminiProcessor(context.Background(), l, cfg.Gemini.APIKey)
	if err != nil {
		l.Fatalf(context.Background(), "Failed to init Gemini Processor: %v", err)
	}
	defer proc.Close()

	// 5. Job Definition
	job := func() {
		ctx := context.Background()
		l.Info(ctx, "Worker: Starting automated crawl job...")

		// A. Crawl
		articles, err := crawlMgr.Run(ctx)
		if err != nil {
			l.Errorf(ctx, "Worker: Crawl failed: %v", err)
			return
		}

		l.Infof(ctx, "Worker: Fetched %d articles. Processing...", len(articles))

		// Define scope for the job
		scope := models.Scope{
			UserID: cfg.Bot.UserID,
			Roles:  []string{"admin"}, // or bot
		}

		for _, article := range articles {
			// Check duplicate EARLY to save AI cost
			_, err = pUC.GetOne(ctx, scope, post.GetOneInput{
				Filter: post.Filter{
					SourceURL: article.SourceURL,
				},
			})
			if err == nil {
				l.Infof(ctx, "Worker: Skipping duplicate article (Pre-Check): %s", article.SourceURL)
				continue
			}

			// B. Process (Using Gemini)
			processed, err := proc.Process(ctx, article)
			if err != nil {
				l.Errorf(ctx, "Worker: Process failed for %s: %v", article.Title, err)
				continue
			}

			// Double check duplicate might be redundant but safe

			// C. Create Post
			content := fmt.Sprintf("![Image](%s)\n\n**%s**\n\n%s\n\nNguồn: %s",
				processed.ImageURL,
				processed.TranslatedTitle,
				processed.TranslatedSummary,
				processed.SourceURL,
			)

			_, err = pUC.Create(ctx, scope, post.CreateInput{
				Content:    content,
				Permission: "public",
				SourceURL:  processed.SourceURL,
			})

			if err != nil {
				l.Errorf(ctx, "Worker: Failed to create post: %v", err)
			} else {
				l.Infof(ctx, "Worker: Created post for '%s'", processed.TranslatedTitle)
			}
		}
	}

	// 6. Scheduler
	c := cron.New()
	// Run every 30 minutes: "*/30 * * * *"
	_, err = c.AddFunc("*/30 * * * *", job)
	if err != nil {
		l.Fatalf(context.Background(), "Error adding cron job: %v", err)
	}

	c.Start()
	l.Info(context.Background(), "Worker started. Press Ctrl+C to stop.")

	// Update: Run once immediately for testing
	go job()

	// Wait for signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	c.Stop()
	l.Info(context.Background(), "Worker stopped")
}
