package main

import (
	"fmt"

	"github.com/bekzod003/link-clean/config"
	"github.com/bekzod003/link-clean/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()

	fmt.Printf("cfg: %v\n", cfg)
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

	log := logger.NewLogger("link bot", loggerLevel)
	// log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

}
