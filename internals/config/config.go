package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	Env          string
	DataBase_URL string
	JwtSecret    string
}

func MustLoad() *Config {

	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Errorf("config.load.file: %v", err))
	}

	port := os.Getenv("PORT")
	if port == "" {
		panic("config: Port is required")
	}
	env := os.Getenv("ENV")
	if env == "" {
		panic("config: ENV is required")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		panic("config: DATABASE_URL is required")
	}

	secretKey := os.Getenv("JWT_SECRET")

	if secretKey == "" {
		panic("config: secretKey is required")

	}

	return &Config{
		Port:         port,
		Env:          env,
		DataBase_URL: dbURL,
		JwtSecret:    secretKey,
	}

}
