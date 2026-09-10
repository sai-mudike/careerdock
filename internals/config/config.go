package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port         string
	Env          string
	DataBase_URL string
}

func MustLoad() *Config {

	godotenv.Load("../../.env")

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

	return &Config{
		Port:         port,
		Env:          env,
		DataBase_URL: dbURL,
	}

}
