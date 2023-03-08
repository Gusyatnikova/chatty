package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Pg       PG
		Password Password
		Http     HTTP
		Jwt      JWT
	}

	HTTP struct {
		Port string `env:"HTTP_PORT"`
		Host string `env:"HTTP_HOST"`
	}

	PG struct {
		DbName   string `env:"POSTGRES_DB"`
		User     string `env:"POSTGRES_USER"`
		Password string `env:"POSTGRES_PASSWORD"`
		Host     string `env:"POSTGRES_HOST"`
		Port     int    `env:"POSTGRES_PORT"`
		PoolMax  int    `env:"POSTGRES_POOL_MAX"`
	}

	JWT struct {
		Sign string `env:"JWT_SECRET"`
		TTL  int64  `env:"JWT_TTL_SEC"`
	}

	Password struct {
		Secret      string `env:"PASSWORD_SECRET"`
		Memory      uint32 `env:"PASSWORD_MEMORY"`
		Iterations  uint32 `env:"PASSWORD_ITERATIONS"`
		SaltLength  uint32 `env:"PASSWORD_SALT_LENGTH"`
		KeyLength   uint32 `env:"PASSWORD_KEY_LENGTH"`
		Parallelism uint8  `env:"PASSWORD_PARALLELISM"`
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
