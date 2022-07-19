package main

import (
	"github.com/bekzod003/link-clean/pkg/logger"
)

func main() {
	loggerLevel := logger.LevelDebug

	// switch cfg.Environment {
	// case config.DebugMode:
	// 	loggerLevel = logger.LevelDebug
	// 	gin.SetMode(gin.DebugMode)
	// case config.TestMode:
	// 	loggerLevel = logger.LevelDebug
	// 	gin.SetMode(gin.TestMode)
	// default:
	// 	loggerLevel = logger.LevelInfo
	// 	gin.SetMode(gin.ReleaseMode)
	// }

	log := logger.NewLogger("link bot", loggerLevel)
	// log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer logger.Cleanup(log)

}
