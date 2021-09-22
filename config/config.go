package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port  string
	PgUrl string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &Config{
		Port:  os.Getenv("PORT"),
		PgUrl: os.Getenv("POSTGRES_URL"),
	}
}
