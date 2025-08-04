package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config contains the configuration for the application
type Config struct {
	JWTKey []byte
}

// cfg is private variable to hold the configuration
var cfg *Config

// init() run once to initialize the configuration
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the JWT key from .env file
	jwtKeyStr := os.Getenv("JWT_KEY")
	if jwtKeyStr == "" {
		log.Fatal("JWT_KEY is not set in the .env file")
	}

	// Initialize the configuration
	cfg = &Config{
		JWTKey: []byte(jwtKeyStr),
	}
}

func Get() *Config {
	return cfg
}
