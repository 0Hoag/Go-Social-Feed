package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

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
	"github.com/hoag/go-social-feed/internal/telegram"
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

	// Init Processor - Using Google Translate (FREE, no quota limits)
	// Gemini is disabled due to API key quota restrictions
	// proc, err := processor.NewGeminiProcessor(context.Background(), l, cfg.Gemini.APIKey)
	// if err != nil {
	// 	l.Fatalf(context.Background(), "Failed to init Gemini Processor: %v", err)
	// }
	// defer proc.Close()

	proc := processor.NewSimpleProcessor(l) // Google Translate - FREE & UNLIMITED

	// Init Telegram Bot
	tgBot, err := telegram.NewTelegramClient(cfg.Telegram.BotToken, cfg.Telegram.ChatID, l)
	if err != nil {
		l.Errorf(context.Background(), "Failed to init Telegram Bot: %v", err)
		// Don't fatal, proceed without telegram
	}

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

		// Limit to 10 most recent articles to avoid spam
		if len(articles) > 10 {
			articles = articles[:10]
			l.Infof(ctx, "Worker: Limited to 10 articles to avoid spam")
		}

		// Define scope for the job
		scope := models.Scope{
			UserID: cfg.Bot.UserID,
			Roles:  []string{"admin"}, // or bot
		}

		for _, article := range articles {
			// Rate Limit: Sleep at start of loop (except first? No, simple is fine)
			// Actually sleep at end is better, but we need to ensure it runs.
			// Let's use a closure to handle the "continue" logic cleanly.

			func() {
				// Check duplicate EARLY to save AI cost
				_, err = pUC.GetOne(ctx, scope, post.GetOneInput{
					Filter: post.Filter{
						SourceURL: article.SourceURL,
					},
				})
				if err == nil {
					l.Infof(ctx, "Worker: Skipping duplicate article (Pre-Check): %s", article.SourceURL)
					return
				}

				// Retry Loop for Gemini
				var processed processor.ProcessedContent
				var processErr error
				maxRetries := 3

				for i := 0; i < maxRetries; i++ {
					processed, processErr = proc.Process(ctx, article)
					if processErr == nil {
						break // Success!
					}

					if strings.Contains(processErr.Error(), "429") {
						time.Sleep(60 * time.Second) // Heavy penalty wait
						continue
					} else {
						// Other error, abort
						break
					}
				}

				if processErr != nil {
					l.Errorf(ctx, "Worker: Process failed for %s after retries: %v", article.Title, processErr)
					return
				}

				// C. Create Post
				content := fmt.Sprintf("![Image](%s)\n\n**%s**\n\n%s\n\nNguồn: %s",
					processed.ImageURL,
					processed.TranslatedTitle,
					processed.TranslatedSummary,
					processed.SourceURL,
				)

				_, err = pUC.Create(ctx, scope, post.CreateInput{
					Content:     content,         // AI summary for feed
					FullContent: article.Content, // Full article text for detail page
					Permission:  "public",
					SourceURL:   processed.SourceURL,
				})

				if err != nil {
					l.Errorf(ctx, "Worker: Failed to create post: %v", err)
				} else {
					l.Infof(ctx, "\n"+
						"✅ POST CREATED SUCCESSFULLY\n"+
						"📝 Title   : %s\n"+
						"🔗 Source  : %s\n"+
						"🖼️ Image   : %s\n"+
						"📜 Content : %s\n"+
						"==================================================",
						processed.TranslatedTitle,
						processed.SourceURL,
						processed.ImageURL,
						processed.TranslatedSummary,
					)

					// Send Notification to Telegram
					if tgBot != nil {
						err := tgBot.SendPost(processed.TranslatedTitle, processed.TranslatedSummary, processed.ImageURL, processed.SourceURL)
						if err != nil {
							l.Errorf(ctx, "Worker: Failed to send Telegram: %v", err)
						} else {
							l.Infof(ctx, "📨 Sent to Telegram")
						}
					}
				}
			}()

			// Always sleep after each attempt (6 minutes to avoid spam)
			time.Sleep(740 * time.Second)
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
