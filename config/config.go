package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr                 string
	FacebookClientID     string
	FacebookClientSecret string
	FacebookCallbackURL  string
	DatabaseHost         string
	DatabasePort         string
	DatabaseUser         string
	DatabasePassword     string
	DatabaseName         string
	SessionSecret        string
	AwsKey               string
	AwsSecret            string
	AwsRegion            string
	AwsImageBucket       string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		Addr:                 os.Getenv("ADDR"),
		FacebookClientID:     os.Getenv("FB_CLIENT_ID"),
		FacebookClientSecret: os.Getenv("FB_CLIENT_SECRET"),
		FacebookCallbackURL:  os.Getenv("FB_CALLBACK_URL"),
		DatabaseHost:         os.Getenv("DB_HOST"),
		DatabasePort:         os.Getenv("DB_PORT"),
		DatabaseUser:         os.Getenv("DB_USER"),
		DatabasePassword:     os.Getenv("DB_PASSWORD"),
		DatabaseName:         os.Getenv("DB_NAME"),
		SessionSecret:        os.Getenv("SESSION_SECRET"),
		AwsKey:               os.Getenv("AWS_ACCESS_KEY_ID"),
		AwsSecret:            os.Getenv("AWS_SECRET_ACCESS_KEY"),
		AwsRegion:            os.Getenv("AWS_REGION"),
		AwsImageBucket:       os.Getenv("AWS_IMAGE_BUCKET"),
	}, nil
}
