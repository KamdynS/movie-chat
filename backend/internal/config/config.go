package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL    string
	ServerAddress  string
	JWTSecret      string
	ClerkSecretKey string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		DatabaseURL:    os.Getenv("DATABASE_URL"),
		ServerAddress:  os.Getenv("SERVER_ADDRESS"),
		JWTSecret:      os.Getenv("JWT_SECRET"),
		ClerkSecretKey: os.Getenv("CLERK_SECRET_KEY"),
	}, nil
}
