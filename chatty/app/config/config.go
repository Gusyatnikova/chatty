package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Http HTTP
		Pg   PG
	}

	HTTP struct {
		Port string `env-required:"true" env:"HTTP_PORT"`
		Host string `env-required:"true" env:"HTTP_HOST"`
	}

	PG struct {
		PoolMax  int    `env-required:"true" env:"POSTGRES_POOL_MAX"`
		Host     string `env-required:"true" env:"POSTGRES_HOST"`
		Port     int    `env-required:"true" env:"POSTGRES_PORT"`
		User     string `env-required:"true" env:"POSTGRES_USER"`
		Password string `env-required:"true" env:"POSTGRES_PASSWORD"`
		DbName   string `env-required:"true" env:"POSTGRES_DB"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("err in config.NewConfing.ReadEnv(): %w", err)
	}

	return cfg, nil
}
