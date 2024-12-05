package config

import (
	"github.com/joho/godotenv"
	"os"
	"sync"
)

type EnvConfig struct {
	// Database
	DBHost     string `env:"DB_HOST"`
	DBPort     string `env:"DB_PORT"`
	DBUser     string `env:"DB_USER"`
	DBPassword string `env:"DB_PASSWORD"`
	DBName     string `env:"DB_NAME"`
	LOGLEVEL   string `env:"LOG_LEVEL"`
}

var (
	envConfig *EnvConfig
	once      sync.Once
)

func GetEnvConfig() *EnvConfig {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			panic(err)
		}

		envConfig = &EnvConfig{
			DBHost:     os.Getenv("DB_HOST"),
			DBPort:     os.Getenv("DB_PORT"),
			DBUser:     os.Getenv("DB_USER"),
			DBPassword: os.Getenv("DB_PASSWORD"),
			DBName:     os.Getenv("DB_NAME"),
			LOGLEVEL:   os.Getenv("LOG_LEVEL"),
		}
	})
	return envConfig
}
