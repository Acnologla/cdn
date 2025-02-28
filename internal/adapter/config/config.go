package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	WasabiConfig struct {
		BucketName string
		AccessKey  string
		SecretKey  string
		Endpoint   string
	}

	HTTPConfig struct {
		AdminKey  string
		Port      string
		ServerURL string
		Secure    bool
	}

	Config struct {
		WasabiConfig WasabiConfig
		HTTPConfig   HTTPConfig
		Production   bool
	}
)

func LoadConfig() (*Config, error) {
	isProduction := os.Getenv("PRODUCTION") == "TRUE"
	if !isProduction {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	WasabiConfig := WasabiConfig{
		BucketName: os.Getenv("WASABI_BUCKET_NAME"),
		AccessKey:  os.Getenv("WASABI_ACCESS_KEY"),
		SecretKey:  os.Getenv("WASABI_SECRET_KEY"),
		Endpoint:   os.Getenv("WASABI_BUCKET_ENDPOINT"),
	}
	HTTPConfig := HTTPConfig{
		AdminKey:  os.Getenv("ADMIN_KEY"),
		Port:      os.Getenv("PORT"),
		ServerURL: os.Getenv("SERVER_URL"),
		Secure:    isProduction && os.Getenv("SECURE") == "TRUE",
	}

	return &Config{
		WasabiConfig: WasabiConfig,
		HTTPConfig:   HTTPConfig,
		Production:   isProduction,
	}, nil
}
