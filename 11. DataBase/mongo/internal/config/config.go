package config

import (
	"time"
)

type Config struct {
	dbURL          string        `toml:"db_url"`
	connectTimeout time.Duration `toml:"timeout"`
}

func New() *Config {
	return &Config{
		dbURL:          "mongodb://localhost:27017/mongo-crud",
		connectTimeout: 10,
	}
}
