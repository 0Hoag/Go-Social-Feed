package httpserver

import (
	"github.com/hoag/go-social-feed/config"
	_ "github.com/hoag/go-social-feed/docs"
	prod "github.com/hoag/go-social-feed/internal/delivery/rabbitmq/producer"
	"github.com/hoag/go-social-feed/pkg/jwt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/hoag/go-social-feed/internal/middleware"
	postHTTP "github.com/hoag/go-social-feed/internal/post/delivery/http"
	postMongo "github.com/hoag/go-social-feed/internal/post/repository/mongo"
	postUC "github.com/hoag/go-social-feed/internal/post/usecase"

	followHTTP "github.com/hoag/go-social-feed/internal/follow/delivery/http"
	followMongo "github.com/hoag/go-social-feed/internal/follow/repository/mongo"
	followUC "github.com/hoag/go-social-feed/internal/follow/usecase"

	commentHTTP "github.com/hoag/go-social-feed/internal/comment/delivery/http"
	commentMongo "github.com/hoag/go-social-feed/internal/comment/repository/mongo"
	commentUC "github.com/hoag/go-social-feed/internal/comment/usecase"

	userHTTP "github.com/hoag/go-social-feed/internal/users/delivery/http"
	userMongo "github.com/hoag/go-social-feed/internal/users/repository/mongo"
	userUC "github.com/hoag/go-social-feed/internal/users/usecase"

	authHTTP "github.com/hoag/go-social-feed/internal/auth/delivery/http"
	authUC "github.com/hoag/go-social-feed/internal/auth/usecase"

	// Import this to execute the init function in docs.go which setups the Swagger docs.
	_ "github.com/hoag/go-social-feed/docs"
)

func (srv HTTPServer) mapHandlers() error {
	srv.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	jwtManager := jwt.NewManager(srv.jwtSecretKey)

	cfg, _ := config.Load()

	// Producer
	postProd := prod.New(srv.l, srv.amqpConn)
	if err := postProd.Run(); err != nil {
		return err
	}

	// Repositories
	userRepo := userMongo.New(srv.l, srv.db)
	postRepo := postMongo.New(srv.l, srv.db)
	followRepo := followMongo.New(srv.l, srv.db)
	commentRepo := commentMongo.New(srv.l, srv.db)

	// Usecases
	userUC := userUC.New(srv.l, userRepo)
	postUC := postUC.New(srv.l, postProd, userUC, postRepo)
	followUC := followUC.New(srv.l, userUC, followRepo)
	commentUC := commentUC.New(srv.l, postUC, commentRepo)
	authUC := authUC.New(srv.l, cfg, userUC)

	// Handlers
	userH := userHTTP.New(srv.l, userUC)
	postH := postHTTP.New(srv.l, postUC)
	followH := followHTTP.New(srv.l, followUC)
	commentH := commentHTTP.New(srv.l, commentUC)
	authH := authHTTP.New(srv.l, authUC)

	// Middlewares
	mw := middleware.New(srv.l, userUC, jwtManager, srv.encrypter, srv.internalKey)

	// Public routes
	srv.gin.Use(mw.Locale())
	api := srv.gin.Group("/api/v1")

	newsFeedGroup := api.Group("/news-feed")
	userHTTP.MapRoutes(newsFeedGroup.Group("/user"), userH, mw)
	authHTTP.MapRoutes(newsFeedGroup.Group("/auth"), authH, mw)
	postHTTP.MapRoutes(newsFeedGroup.Group("/posts"), postH, mw)
	followHTTP.MapRoutes(newsFeedGroup.Group("/follow"), followH, mw)
	commentHTTP.MapRoutes(newsFeedGroup.Group("/comment"), commentH, mw)

	return nil
}
