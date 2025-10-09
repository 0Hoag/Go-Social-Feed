package httpserver

import (
	"github.com/hoag/go-social-feed/config"
	_ "github.com/hoag/go-social-feed/docs"
	"github.com/hoag/go-social-feed/pkg/jwt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/hoag/go-social-feed/internal/middleware"
	postHTTP "github.com/hoag/go-social-feed/internal/post/delivery/http"
	postMongo "github.com/hoag/go-social-feed/internal/post/repository/mongo"
	postUC "github.com/hoag/go-social-feed/internal/post/usecase"

	reactionHTTP "github.com/hoag/go-social-feed/internal/reaction/delivery/http"
	reactionMongo "github.com/hoag/go-social-feed/internal/reaction/repository/mongo"
	reactionUC "github.com/hoag/go-social-feed/internal/reaction/usecase"

	followHTTP "github.com/hoag/go-social-feed/internal/follow/delivery/http"
	followMongo "github.com/hoag/go-social-feed/internal/follow/repository/mongo"
	followUC "github.com/hoag/go-social-feed/internal/follow/usecase"

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

	// Repositories
	postRepo := postMongo.New(srv.l, srv.db)
	reactionRepo := reactionMongo.New(srv.l, srv.db)
	followRepo := followMongo.New(srv.l, srv.db)
	userRepo := userMongo.New(srv.l, srv.db)

	// Usecases
	postUC := postUC.New(srv.l, postRepo)
	reactionUC := reactionUC.New(srv.l, postUC, reactionRepo)
	userUC := userUC.New(srv.l, userRepo)
	followUC := followUC.New(srv.l, userUC, followRepo)
	authUC := authUC.New(srv.l, cfg, userUC)

	// Handlers
	postH := postHTTP.New(srv.l, postUC)
	reactionH := reactionHTTP.New(srv.l, reactionUC)
	followH := followHTTP.New(srv.l, followUC)
	userH := userHTTP.New(srv.l, userUC)
	authH := authHTTP.New(srv.l, authUC)

	// Middlewares
	mw := middleware.New(srv.l, userUC, jwtManager, srv.encrypter, srv.internalKey)

	// Public routes
	srv.gin.Use(mw.Locale())
	api := srv.gin.Group("/api/v1")

	// Routes
	newsFeedGroup := api.Group("/news-feed")
	authHTTP.MapRoutes(newsFeedGroup.Group("/auth"), authH, mw)
	postHTTP.MapRoutes(newsFeedGroup.Group("/posts"), postH, mw)
	reactionHTTP.MapRoutes(newsFeedGroup.Group("/reaction"), reactionH, mw)
	followHTTP.MapRoutes(newsFeedGroup.Group("/follow"), followH, mw)
	userHTTP.MapRoutes(newsFeedGroup.Group("/user"), userH, mw)

	return nil
}
