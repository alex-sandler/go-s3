package config

import (
	"fmt"

	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	GRPC GRPCConfig `envPrefix:"GRPC_" env-required:"true"`
}

type GRPCConfig struct {
	Host string `env:"HOST" env-required:"true"`
	Port string `env:"PORT" env-required:"true"`
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
