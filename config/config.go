package config

import "github.com/caarlos0/env/v9"

// Config represents the application configuration
type Config struct {
	Logger       LoggerConfig
	HTTPServer   HTTPServerConfig
	Mongo        MongoConfig
	RabbitConfig RabbitConfig
	JWT          JWTConfig
	Encrypter    EncrypterConfig
	Bot          BotConfig
	Gemini       GeminiConfig
	Telegram     TelegramConfig
	Scanner      ScannerConfig
}

type ScannerConfig struct {
	EtherscanAPIKey    string `env:"ETHERSCAN_API_KEY"`
	BscScanAPIKey      string `env:"BSCSCAN_API_KEY"`
	BaseScanAPIKey     string `env:"BASESCAN_API_KEY"`
	ArbitrumScanAPIKey string `env:"ARBITRUMSCAN_API_KEY"`
	PolygonScanAPIKey  string `env:"POLYGONSCAN_API_KEY"`
}

type GeminiConfig struct {
	APIKey string `env:"GEMINI_API_KEY"`
}

type TelegramConfig struct {
	BotToken string `env:"TELEGRAM_BOT_TOKEN"`
	ChatID   int64  `env:"TELEGRAM_CHAT_ID"`
}

type BotConfig struct {
	UserID string `env:"BOT_USER_ID"`
}

type JWTConfig struct {
	SecretKey string `env:"JWT_SECRET"`
}

type RabbitConfig struct {
	URL string `env:"RABBITMQ_URL"`
}

type HTTPServerConfig struct {
	Port int    `env:"APP_PORT" envDefault:"80"`
	Mode string `env:"API_MODE" envDefault:"debug"`
}

type LoggerConfig struct {
	Level    string `env:"LOGGER_LEVEL" envDefault:"debug"`
	Mode     string `env:"LOGGER_MODE" envDefault:"development"`
	Encoding string `env:"LOGGER_ENCODING" envDefault:"console"`
}

type MongoConfig struct {
	Database       string `env:"MONGODB_DATABASE"`
	URI            string `env:"MONGODB_URI"`
	ENABLE_MONITOR bool   `env:"MONGODB_ENABLE_MONITORING" envDefault:"false"`
}

type EncrypterConfig struct {
	Key string `env:"ENCRYPT_KEY"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
