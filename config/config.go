package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Schema   string
}

type MapConfig struct {
	APIKey string
}

type JWTConfig struct {
	Secret string
}

type S3Config struct {
	Access   string
	Endpoint string
	Secret   string
	Bucket   string
}

type Config struct {
	IsProduction bool
	DB           *DatabaseConfig
	Map          *MapConfig
	JWT          *JWTConfig
	S3           *S3Config
}

var cfg *Config

// loadString retrieves a string environment variable or logs a fatal error if it's missing.
func loadString(key string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Fatal("Environment variable " + key + " is not set")
	}
	return val
}

// loadInt retrieves an integer environment variable or logs a fatal error if it's invalid.
func loadInt(key string) int {
	val := loadString(key)
	intVal, err := strconv.Atoi(val)
	if err != nil {
		log.Fatal("Environment variable " + key + " is not a valid number")
	}
	return intVal
}

// MustLoad initializes the configuration by loading environment variables.
func MustLoad() {
	env := os.Getenv("ENV")

	isProduction := strings.ToLower(env) == "production"
	if !isProduction {
		_ = godotenv.Load() // Safely load .env without worrying about errors
	}

	cfg = &Config{
		IsProduction: isProduction,
		DB: &DatabaseConfig{
			Host:     loadString("DB_HOST"),
			Port:     loadInt("DB_PORT"),
			Name:     loadString("DB_NAME"),
			User:     loadString("DB_USER"),
			Password: loadString("DB_PASSWORD"),
			Schema:   loadString("DB_SCHEMA"),
		},
		JWT: &JWTConfig{
			Secret: loadString("JWT_SECRET"),
		},
	}
}

// Get returns the loaded configuration, ensuring it was initialized first.
func Get() *Config {
	if cfg == nil {
		log.Fatal("Configuration not loaded. Call MustLoad before using Get.")
	}
	return cfg
}
