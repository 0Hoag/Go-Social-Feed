package httpserver

import (
	"github.com/gin-gonic/gin"
	pkgCrt "github.com/hoag/go-social-feed/pkg/encrypter"
	pkgLog "github.com/hoag/go-social-feed/pkg/log"
	pkgMongo "github.com/hoag/go-social-feed/pkg/mongo"
	"github.com/hoag/go-social-feed/pkg/rabbitmq"
)

const productionMode = "production"

var ginMode = gin.DebugMode

type HTTPServer struct {
	gin          *gin.Engine
	l            pkgLog.Logger
	port         int
	db           pkgMongo.Database
	amqpConn     rabbitmq.Connection
	jwtSecretKey string
	mode         string
	hoagConfig   HoagConfig
	internalKey  string
	encrypter    pkgCrt.Encrypter
	secretConfig SecretConfig
}

type Config struct {
	Port         int
	JWTSecretKey string
	DB           pkgMongo.Database
	AMQPConn     rabbitmq.Connection
	Mode         string
	HoagConfig   HoagConfig
	InternalKey  string
	Encrypter    pkgCrt.Encrypter
	SecretConfig SecretConfig
}

type HoagConfig struct {
	AdminDomain string
}

type SecretConfig struct {
	SecretKey string
}

func New(l pkgLog.Logger, cfg Config) *HTTPServer {
	if cfg.Mode == productionMode {
		ginMode = gin.ReleaseMode
	}

	gin.SetMode(ginMode)

	return &HTTPServer{
		l:            l,
		gin:          gin.Default(),
		port:         cfg.Port,
		db:           cfg.DB,
		amqpConn:     cfg.AMQPConn,
		jwtSecretKey: cfg.JWTSecretKey,
		mode:         cfg.Mode,
		hoagConfig:   cfg.HoagConfig,
		internalKey:  cfg.InternalKey,
		encrypter:    cfg.Encrypter,
		secretConfig: cfg.SecretConfig,
	}
}
