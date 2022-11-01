package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
		PG   `yaml:"postgres"`
	}

	HTTP struct {
		Port string `env-required:"true" env:"HTTP_PORT"`
		Host string `env-required:"true" env:"HTTP_HOST"`
	}

	Log struct {
		Level string `env-required:"true" env:"LOG_LEVEL"`
	}

	PG struct {
		PoolMax  int    `env-required:"true" env:"PG_POOL_MAX"`
		Host     string `env-required:"true" env:"PG_HOST"`
		Port     int    `env-required:"true" env:"PG_PORT"`
		User     string `env-required:"true" env:"PG_USER"`
		Password string `env-required:"true" env:"PG_PASSWORD"`
		DbName   string `env-required:"true" env:"PG_DB_NAME"`
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
