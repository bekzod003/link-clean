package config

import (
	"log"

	"github.com/Netflix/go-env"
)

type Config struct {
	HTTP struct {
		Port string `env:"PORT,required=true"`
	}
	ServiceName string `env:"SERVICE_NAME"`
	Environment string `env:"ENVIRONMENT,default:development"`

	PostgreSQL struct {
		Host     string `env:"POSTGRES_HOST,required=true"`
		Port     int    `env:"POSTGRES_PORT,required=true"`
		User     string `env:"POSTGRES_USER,required=true"`
		Password string `env:"POSTGRES_PASSWORD,required=true"`
		DBName   string `env:"POSTGRES_DB_NAME,required=true"`
	}
}

func GetConfig() *Config {
	cfg := Config{}
	if _, err := env.UnmarshalFromEnviron(&cfg); err != nil {
		log.Fatal("Error while getting config: ", err)
		return nil
	}

	return &cfg
}
