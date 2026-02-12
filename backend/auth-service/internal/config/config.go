package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT      string
	LoggerEnv string
	DSN       string
}

var Envs = Init()

func Init() *Config {
	projectRoot := getProjectRoot()
	envPath := filepath.Join(projectRoot, ".env")

	if err := godotenv.Load(envPath); err != nil {
		log.Printf("No .env file found at %s, using environment variables\n", envPath)
	} else {
		log.Printf("Successfully loaded .env from %s\n", envPath)
	}

	loggerEnv := getOrDefault("APP_ENV_SERVICE", "development")
	port := getOrDefault("PORT_AUTH", "8080")
	dbHost := getOrDefault("DB_HOST", "localhost")
	dbPort := getOrDefault("DB_PORT", "5432")
	dbUser := getOrDefault("DB_USER", "postgres")
	dbPassword := getOrDefault("DB_PASSWORD", "postgres")
	dbName := getOrDefault("DB_NAME", "postgres")

	return &Config{
		LoggerEnv: loggerEnv,
		PORT:      port,
		DSN: fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			dbHost, dbPort, dbUser, dbPassword, dbName),
	}
}

func getOrDefault(varName string, defaultValue string) string {
	value := os.Getenv(varName)
	if value == "" {
		return defaultValue
	}

	return value
}

func getProjectRoot() string {
	_, filename, _, _ := runtime.Caller(0)

	projectRoot := filepath.Join(filepath.Dir(filename), "..", "..", "..")
	return projectRoot
}
