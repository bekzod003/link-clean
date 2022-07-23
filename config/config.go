package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	HTTP struct {
		Port string `env:"PORT" env-required:"true"`
	}
	ServiceName string `env:"SERVICE_NAME"`
	Environment string `env:"ENVIRONMENT" env-default:"development"`

	PostgreSQL struct {
		Host     string `env:"POSTGRES_HOST" env-required:"true"`
		Port     int    `env:"POSTGRES_PORT" env-required:"true"`
		User     string `env:"POSTGRES_USER" env-required:"true"`
		Password string `env:"POSTGRES_PASSWORD" env-required:"true"`
		DBName   string `env:"POSTGRES_DB_NAME" env-required:"true"`

		PoolConfig struct {
			MaxConns                int32 `env:"POSTGRES_MAX_CONNS" env-default:"30"`
			MaxConnIdleMinute       int   `env:"POSTGRES_MAX_CONN_IDLE_MINUTES" env-default:"40"`
			MaxConnLifetimeMinute   int   `env:"POSTGRES_CONN_LIFE_MINUTES" env-default:"30"`
			HealthCheckPeriodMinute int   `env:"POSTGRES_HEALTH_CHECK_PERIOD_MINUTES" env-default:"10"`
		}
	}
}

var (
	onceGetConfig sync.Once
	cfg           *Config
)

func GetConfig() *Config {
	onceGetConfig.Do(func() {
		cfg = getConfig()
	})
	return cfg
}

func getConfig() *Config {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfgTemp := Config{}
	if err = cleanenv.ReadEnv(&cfgTemp); err != nil {
		log.Fatal("Error while getting config: ", err)
		return nil
	}

	return &cfgTemp
}
