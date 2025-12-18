package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hoag/go-social-feed/config"
	pkgCrt "github.com/hoag/go-social-feed/pkg/encrypter"
	"github.com/hoag/go-social-feed/pkg/mongo"
)

const (
	connectTimeout = 10 * time.Second
)

func Connect(cfg config.MongoConfig, encrypter pkgCrt.Encrypter) (mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeout)
	defer cancel()

	if cfg.URI == "" {
		return nil, errors.New("mongo uri is empty")
	}

	uri := cfg.URI

	// 🔐 CHỈ decrypt nếu KHÔNG phải Atlas
	if encrypter != nil && !strings.HasPrefix(uri, "mongodb+srv://") {
		if dec, err := encrypter.Decrypt(uri); err == nil && dec != "" {
			uri = dec
		} else {
			log.Printf("warning: mongo uri decrypt failed, using plaintext: %v", err)
		}
	}

	client, err := mongo.Connect(ctx, mongo.NewClientOptions().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Atlas bắt buộc Primary
	err = client.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping to media DB: %w", err)
	}

	log.Println("✅ Connected to MongoDB")
	return client, nil
}

// Disconnect disconnects from the database.
func Disconnect(mediaClient mongo.Client) {
	if mediaClient == nil {
		return
	}

	err := mediaClient.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
