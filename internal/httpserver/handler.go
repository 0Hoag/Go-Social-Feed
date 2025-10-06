package httpserver

import (
	_ "github.com/hoag/go-social-feed/docs"
	"github.com/hoag/go-social-feed/pkg/jwt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/hoag/go-social-feed/internal/middleware"
	postHTTP "github.com/hoag/go-social-feed/internal/post/delivery/http"
	postMongo "github.com/hoag/go-social-feed/internal/post/repository/mongo"
	postUC "github.com/hoag/go-social-feed/internal/post/usecase"

	userHTTP "github.com/hoag/go-social-feed/internal/users/delivery/http"
	userMongo "github.com/hoag/go-social-feed/internal/users/repository/mongo"
	userUC "github.com/hoag/go-social-feed/internal/users/usecase"
)

func (srv HTTPServer) mapHandlers() error {
	srv.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	jwtManager := jwt.NewManager(srv.jwtSecretKey)

	// Repositories
	postRepo := postMongo.New(srv.l, srv.db)
	userRepo := userMongo.New(srv.l, srv.db)

	// Usecases
	postUC := postUC.New(srv.l, postRepo)
	userUC := userUC.New(srv.l, userRepo)

	// Handlers
	postH := postHTTP.New(srv.l, postUC)
	userH := userHTTP.New(srv.l, userUC)

	// Middlewares
	mw := middleware.New(srv.l, userUC, jwtManager, srv.encrypter, srv.internalKey)

	// Public routes
	srv.gin.Use(mw.Locale())
	api := srv.gin.Group("/api/v1")

	// Routes

	return nil
}
