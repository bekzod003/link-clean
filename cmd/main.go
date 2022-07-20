package main

import (
	"github.com/bekzod003/link-clean/config"
	"github.com/bekzod003/link-clean/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	println("Config intializing...")
	cfg := config.GetConfig()
	println("Config intialized successfully")

	loggerLevel := logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	println("Logger intializing...")
	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	println("Logger intialized successfully")

	defer logger.Cleanup(log)

	log.Info("Starting link bot")
}
