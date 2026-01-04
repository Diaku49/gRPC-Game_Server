package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	//Server
	Port      string `envconfig:"PORT" default:"3000"`
	JwtSecret string `envconfig:"JWT_SECRET" default:"somethingeasy"`
	//DB
	DbUrl string `envconfig:"DB_URL" default:"postgres://postgres:postgres@localhost:5432/postgres"`
	//Redis
	RedisAddr string `envconfig:"REDIS_ADDR" default:"localhost:6379"`
	RedisPass string `envconfig:"REDIS_PASS" default:""`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("failed to process environment variables: %w", err)
	}
	return &cfg, nil
}
