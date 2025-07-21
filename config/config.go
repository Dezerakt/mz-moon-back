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
	}

	App struct {
		Port int    `env:"APP_PORT"`
		Name string `env:"APP_NAME"`
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

func InitConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}

func GetEnvVar() {
}
