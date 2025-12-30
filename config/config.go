package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port        string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	} else {
		log.Println(".env file loaded successfully")
	}

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://localhost:5432/mydb"),
		Port:        getEnv("PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		log.Printf("Found env var %s: %s", key, value)
		return value
	}
	log.Printf("Using default for %s: %s", key, defaultValue)
	return defaultValue
}
