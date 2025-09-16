package config

import "github.com/caarlos0/env/v9"

type Config struct {
	HTTPServer HTTPServerConfig
	Logger     LoggerConfig
	JWT        JWTConfig
	Mongo      MongoConfig
	Encrypter  EncrypterConfig
}

type JWTConfig struct {
	SecretKey string `env:"JWT_SECRET"`
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
	Database            string `env:"MONGODB_DATABASE"`
	MONGODB_ENCODED_URI string `env:"MONGODB_ENCODED_URI"`
	ENABLE_MONITOR      bool   `env:"MONGODB_ENABLE_MONITORING" envDefault:"false"`
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
