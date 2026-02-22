package httpserver

import (
	"context"
	"fmt"

	"github.com/hoag/go-social-feed/internal/seeder"
)

func (srv HTTPServer) Run() error {
	err := srv.mapHandlers()
	if err != nil {
		return err
	}

	ctx := context.Background()

	// Seed database on startup (idempotent — skips existing records)
	if err := seeder.Run(ctx, srv.db); err != nil {
		srv.l.Errorf(ctx, "seeder.Run: %v", err)
		// Non-fatal: log and continue
	}

	srv.l.Infof(ctx, "Started server on :%d", srv.port)
	return srv.gin.Run(fmt.Sprintf(":%d", srv.port))
}
