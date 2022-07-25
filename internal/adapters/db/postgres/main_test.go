package postgres

import (
	"context"
	"log"
	"time"

	"github.com/bekzod003/link-clean/config"
	"github.com/bekzod003/link-clean/pkg/database/client/postgresql"
)

var (
	deafultTimeOut = time.Second * 10
)

func newClient() postgresql.Client {
	cfg := config.GetConfig()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	client, err := postgresql.NewClient(
		ctx, postgresql.ClientConfig{
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
		},
	)
	if err != nil {
		log.Fatal("Error while getting new postgres client")
	}

	return client
}
