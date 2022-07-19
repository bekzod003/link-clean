package main

import (
	"fmt"

	"github.com/bekzod003/link-clean/config"
	"github.com/bekzod003/link-clean/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()
	cfg1 := config.GetConfig()
	cfg2 := config.GetConfig()
	cfg3 := config.GetConfig()
	cfg4 := config.GetConfig()
	cfg5 := config.GetConfig()

	fmt.Printf("cfg: %v\n", cfg1)
	fmt.Printf("cfg: %v\n", cfg2)
	fmt.Printf("cfg: %v\n", cfg3)
	fmt.Printf("cfg: %v\n", cfg4)
	fmt.Printf("cfg: %v\n", cfg5)
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

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

	log.Info("Starting link bot")
}
