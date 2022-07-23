package main

import (
	"github.com/bekzod003/link-clean/config"
	"github.com/bekzod003/link-clean/internal/app"
)

func main() {
	println("Config intializing...")
	cfg := config.GetConfig()
	println("Config intialized successfully")

	app.Run(cfg)
}
