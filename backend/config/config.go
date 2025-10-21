package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
	ServerPort string
	Env        string
}

var AppConfig *Config

// LoadConfig membaca .env file dan set config
func LoadConfig() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️  Warning: .env file not found, using environment variables")
	}

	AppConfig = &Config{
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", "sk8consign"),
		JWTSecret:  getEnv("JWT_SECRET", "your-secret-key"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
		Env:        getEnv("ENV", "development"),
	}

	log.Println("✅ Configuration loaded")
	log.Printf("   Database: %s@%s:%s/%s", AppConfig.DBUser, AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBName)
	log.Printf("   Environment: %s", AppConfig.Env)
}

// getEnv helper untuk ambil env dengan default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
