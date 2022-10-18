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
		Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
		Host string `env-required:"true" yaml:"host" env:"HTTP_HOST"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"`
	}

	PG struct {
		PoolMax  int    `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		Host     string `env-required:"true" yaml:"host" env:"PG_HOST"`
		Port     int    `env-required:"true" yaml:"port" env:"PG_PORT"`
		User     string `env-required:"true" env:"PG_USER"`
		Password string `env-required:"true" env:"PG_PASSWORD"`
		DbName   string `env-required:"true" env:"PG_DB_NAME"`
	}
)

//todo:  all envs move to environment
//see viper?
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./chatty/app/config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("err in config.NewConfing.ReadConfig(): %w", err)
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("err in config.NewConfing.ReadEnv(): %w", err)
	}

	return cfg, nil
}
