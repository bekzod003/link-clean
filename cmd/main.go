package main

import (
	"github.com/bekzod003/link-clean/config"
	"github.com/bekzod003/link-clean/internal/app"
)

func main() {
	println("Config initializing...")
	cfg := config.GetConfig()
	println("Config initialized successfully")

	app.Run(cfg)
}
