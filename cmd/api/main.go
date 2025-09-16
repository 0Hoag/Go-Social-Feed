package main

import (
	"github.com/hoag/go-social-feed/config"
	httpserver "github.com/hoag/go-social-feed/internal/httpserver"

	"github.com/hoag/go-social-feed/internal/appconfig/mongo"
	pkgCrt "github.com/hoag/go-social-feed/pkg/encrypter"
	pkgLog "github.com/hoag/go-social-feed/pkg/log"
)

func main() {
	// Load config
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

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

	srv := httpserver.New(l, httpserver.Config{
		Port:         cfg.HTTPServer.Port,
		DB:           db,
		JWTSecretKey: cfg.JWT.SecretKey,
		Mode:         cfg.HTTPServer.Mode,
		Encrypter:    crp,
		SecretConfig: httpserver.SecretConfig{
			SecretKey: cfg.Encrypter.Key,
		},
	})
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
