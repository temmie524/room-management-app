package config

import "os"

type AppConfig struct {
	SECRET_KEY string
	API_DOMAIN string
	FE_URL     string
}

func NewAppConfig() *AppConfig {
	cnf := AppConfig{
		SECRET_KEY: os.Getenv("SECRET"),
		API_DOMAIN: os.Getenv("API_DOMAIN"),
		FE_URL:     os.Getenv("FE_URL"),
	}
	return &cnf
}
