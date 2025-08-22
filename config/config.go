package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App      App
		Postgres Postgres
		Mongo    Mongo
		Http     Http
	}
	Http struct {
		TrustedOrigins []string `env:"HTTP_TRUSTED_ORIGINS" envDefault:"https://localhost"`
	}

	App struct {
		Port    int    `env:"APP_PORT"`
		Name    string `env:"APP_NAME"`
		Version string `env:"APP_VERSION"`
	}

	Mongo struct {
		Host     string `env:"MONGO_HOST"`
		Port     int    `env:"MONGO_PORT"`
		User     string `env:"MONGO_USER"`
		Password string `env:"MONGO_PASSWORD"`
	}

	Postgres struct {
		Host     string `env:"POSTGRES_HOST"`
		Port     string `env:"POSTGRES_PORT"`
		Db       string `env:"POSTGRES_DB"`
		Password string `env:"POSTGRES_PASSWORD"`
		User     string `env:"POSTGRES_USER"`
	}
)

var config *Config

func InitConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	config = cfg

	return cfg, nil
}

func GetAppVersion() string {
	return config.App.Version
}

func GetTrustedOrigins() []string { return config.Http.TrustedOrigins }
