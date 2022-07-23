package app

import (
	"github.com/bekzod003/link-clean/config"
	"github.com/bekzod003/link-clean/pkg/logger"
)

func Run(cfg *config.Config) {
	println("Logger intializing...")
	log := logger.NewLogger(cfg.ServiceName, cfg.LoggerLevel)
	println("Logger intialized successfully")

	defer logger.Cleanup(log)

}
