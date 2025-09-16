package mongo

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hoag/go-social-feed/config"
	pkgCrt "github.com/hoag/go-social-feed/pkg/encrypter"
	"github.com/hoag/go-social-feed/pkg/mongo"
)

const (
	connectTimeout = 10 * time.Second
)

// Connect connects to the database
func Connect(mongoConfig config.MongoConfig, encrypter pkgCrt.Encrypter) (mongo.Client, error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), connectTimeout)
	defer cancelFunc()

	rawURI := mongoConfig.MONGODB_ENCODED_URI
	if rawURI == "" {
		return nil, fmt.Errorf("mongo uri is empty")
	}

	uri := rawURI
	if encrypter != nil {
		if dec, err := encrypter.Decrypt(rawURI); err == nil && dec != "" {
			uri = dec
		} else {
			log.Printf("warning: using plaintext Mongo URI (decrypt failed: %v)", err)
		}
	}

	opts := mongo.NewClientOptions().ApplyURI(uri)

	if mongoConfig.ENABLE_MONITOR {
		opts.SetMonitor(mongo.CommandMonitor{
			Started: func(ctx context.Context, e *mongo.CommandStartedEvent) {
				log.Printf("MongoDB command started: %v", e.Command)
			},
			Succeeded: func(ctx context.Context, e *mongo.CommandSucceededEvent) {
				log.Printf("MongoDB command succeeded: %v", e.Reply)
			},
			Failure: func(ctx context.Context, e *mongo.CommandFailedEvent) {
				log.Printf("MongoDB command failed: %v", e.Failure)
			},
		})
	}

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to media DB: %w", err)
	}

	err = client.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping to media DB: %w", err)
	}

	log.Println("Connected to MongoDB!")

	return client, nil
}

// Disconnect disconnects from the database.
func Disconnect(mediaClient mongo.Client) {
	if mediaClient == nil {
		return
	}

	err := mediaClient.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
