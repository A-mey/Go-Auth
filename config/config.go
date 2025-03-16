package config

import (
	"fmt"
	"log"
	"os"

	"github.com/A-mey/GO-AUTH/common/models"
	"github.com/joho/godotenv"
)

func LoadConfig() models.Config {

	env := os.Getenv("APP_ENV")

	envFile := ".env"

	fmt.Println("environment1", env)

	if env == "development" {
		envFile = ".env.dev"
	} else if env == "production" {
		envFile = ".env.prod"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("No %s file found. Using system environment variables.\n", envFile)
	}

	config := models.Config{
		AppEnv:     getEnv("APP_ENV"),
		DBHost:     getEnv("DB_HOST"),
		DBPort:     getEnv("DB_PORT"),
		DBUser:     getEnv("DB_USER"),
		DBPassword: getEnv("DB_PASSWORD"),
		DBName:     getEnv("DB_NAME"),
	}

	fmt.Println("config", config)

	return config
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	} else {
		log.Fatalf("Environment variable %s is not set", key)
		return ""
	}
}
