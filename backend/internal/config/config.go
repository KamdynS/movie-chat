package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress  string
	ClerkSecretKey string
	ClerkPublicKey string
}

func Load() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	return &Config{
		ServerAddress:  os.Getenv("SERVER_ADDRESS"),
		ClerkSecretKey: os.Getenv("CLERK_SECRET_KEY"),
		ClerkPublicKey: os.Getenv("CLERK_PUBLIC_KEY"),
	}, nil
}
