package main

import (
	"flag"
	"log"
	"mongo-crud/internal/config"

	"github.com/BurntSushi/toml"
)

var (
	cfgPath string
)

func init() {
	flag.StringVar(&cfgPath, "config-path", "config/config.toml", "path to config path")
}

func main() {
	flag.Parse()

	config := config.New()

	_, err := toml.DecodeFile(cfgPath, config)
	if err != nil {
		log.Fatalf("can't decode config: %s", err)
	}

}
