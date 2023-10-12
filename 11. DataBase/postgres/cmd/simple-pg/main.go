package main

import (
	"github.com/yellowpuki/simple-pg/internal/config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

	// TODO: Config
	cfg := config.MustLoad()

	// TODO: Logger
	// TODO: Storage
	// TODO: Router
	// TODO: Server
}
