package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTP  HTTPConfig  `envPrefix:"HTTP_" env-required:"true"`
	Minio MinioConfig `envPrefix:"MINIO_" env-required:"true"`
	Image ImageConfig `envPrefix:"IMAGE_" env-required:"true"`
}

type HTTPConfig struct {
	Host string `env:"HOST" env-required:"true"`
	Port string `env:"PORT" env-required:"true"`
}

type MinioConfig struct {
	Endpoint           string `env:"ENDPOINT" env-required:"true"`
	RootUser           string `env:"ROOT_USER" env-required:"true"`
	RootPassword       string `env:"ROOT_PASSWORD" env-required:"true"`
	BucketName         string `env:"BUCKET_NAME" env-required:"true"`
	UseSSL             bool   `env:"USE_SSL" env-required:"true"`
	FileTimeExpiration int    `env:"FILE_TIME_EXPIRATION" env-required:"true"`
}

type ImageConfig struct {
	Quality  int `env:"QUALITY" env-required:"true"`
	MaxWidth int `env:"MAX_WIDTH" env-required:"true"`
}

func Load() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("config.Load: no .env file found: %w", err)
	}

	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("config.Load: failed to parse config: %w", err)
	}

	return &cfg, nil
}
