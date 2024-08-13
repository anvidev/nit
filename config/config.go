package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr                 string
	FacebookClientID     string
	FacebookClientSecret string
	DatabaseHost         string
	DatabasePort         string
	DatabaseUser         string
	DatabasePassword     string
	DatabaseName         string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error loading .env file: %v\n", err)
	}

	return &Config{
		Addr:                 os.Getenv("ADDR"),
		FacebookClientID:     os.Getenv("FB_CLIENT_ID"),
		FacebookClientSecret: os.Getenv("FB_CLIENT_SECRET"),
		DatabaseHost:         os.Getenv("DB_HOST"),
		DatabasePort:         os.Getenv("DB_PORT"),
		DatabaseUser:         os.Getenv("DB_USER"),
		DatabasePassword:     os.Getenv("DB_PASSWORD"),
		DatabaseName:         os.Getenv("DB_NAME"),
	}
}
