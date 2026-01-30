package main

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/hoag/go-social-feed/config"
	httpserver "github.com/hoag/go-social-feed/internal/httpserver"

	"github.com/hoag/go-social-feed/internal/appconfig/mongo"
	pkgCrt "github.com/hoag/go-social-feed/pkg/encrypter"
	pkgLog "github.com/hoag/go-social-feed/pkg/log"
	"github.com/hoag/go-social-feed/pkg/rabbitmq"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("DEBUG: main started")

	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found, using environment variables")
	}

	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	fmt.Println("DEBUG: config loaded")

	crp := pkgCrt.NewEncrypter(cfg.Encrypter.Key)
	client, err := mongo.Connect(cfg.Mongo, crp)
	if err != nil {
		panic(err)
	}
	defer mongo.Disconnect(client)

	db := client.Database(cfg.Mongo.Database)

	l := pkgLog.InitializeZapLogger(pkgLog.ZapConfig{
		Level:    cfg.Logger.Level,
		Mode:     cfg.Logger.Mode,
		Encoding: cfg.Logger.Encoding,
	})

	amqpConn, err := rabbitmq.Dial(cfg.RabbitConfig.URL, true)
	if err != nil {
		l.Warnf(context.Background(), "RabbitMQ not connected, running without queue...")
		amqpConn = rabbitmq.Connection{}
	}
	defer amqpConn.Close()

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = strconv.Itoa(cfg.HTTPServer.Port) // fallback when running locally
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}

	srv := httpserver.New(l, httpserver.Config{
		Port:         port,
		DB:           db,
		AMQPConn:     amqpConn,
		JWTSecretKey: cfg.JWT.SecretKey,
		Mode:         cfg.HTTPServer.Mode,
		Encrypter:    crp,
		SecretConfig: httpserver.SecretConfig{
			SecretKey: cfg.Encrypter.Key,
		},
	})
	fmt.Println("DEBUG: before srv.Run()")
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
