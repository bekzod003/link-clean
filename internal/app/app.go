package app

import (
	"context"
	"github.com/bekzod003/link-clean/config"
	"github.com/bekzod003/link-clean/pkg/database/client/postgresql"
	"github.com/bekzod003/link-clean/pkg/logger"
	"go.uber.org/zap"
	"time"
)

func Run(cfg *config.Config) {
	println("Logger initializing...")
	log := logger.NewLogger(cfg.ServiceName, cfg.LoggerLevel)
	println("Logger initialized successfully")
	defer logger.Cleanup(log)

	psqlConfig := getPostgresClientConfig(cfg)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	poolConnection, err := postgresql.NewClient(ctx, psqlConfig)
	if err != nil {
		log.Fatal("Error while getting postgres new client", zap.Error(err))
	}
	defer poolConnection.Close()

}

func getPostgresClientConfig(cfg *config.Config) postgresql.ClientConfig {
	return postgresql.ClientConfig{
		Login:    cfg.PostgreSQL.User,
		Password: cfg.PostgreSQL.Password,
		Host:     cfg.PostgreSQL.Host,
		Port:     cfg.PostgreSQL.Port,
		DBName:   cfg.PostgreSQL.DBName,
		PoolConfig: postgresql.PoolConfig{
			MaxConns:                 cfg.PostgreSQL.PoolConfig.MaxConns,
			MaxConnIdleMinutes:       cfg.PostgreSQL.PoolConfig.MaxConnIdleMinute,
			MaxConnLifetimeMinutes:   cfg.PostgreSQL.PoolConfig.MaxConnLifetimeMinute,
			HealthCheckPeriodMinutes: cfg.PostgreSQL.PoolConfig.HealthCheckPeriodMinute,
		},
	}
}
