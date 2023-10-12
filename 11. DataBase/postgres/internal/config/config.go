package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	ENV         string `yaml:"env" envDefault:"local"`
	StoragePath string `yaml:"storage_path" envRequired:"true"`
	HTTPServer  `yaml:"http_server" `
}

type HTTPServer struct {
	Address     string        `yaml:"address" envDefault:"localhost:8081"`
	Timeout     time.Duration `yaml:"timeout" envDefault:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" envDefault:"60s"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file %s not exist", configPath)
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatal("can't read config")
	}

	return &cfg
}
