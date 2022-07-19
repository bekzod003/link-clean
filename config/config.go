package config

import (
	"log"

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
	}
}

func GetConfig() *Config {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{}
	if err = cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatal("Error while getting config: ", err)
		return nil
	}

	return &cfg
}
