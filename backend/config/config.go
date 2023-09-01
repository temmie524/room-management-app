package config

import "os"

type AppConfig struct {
	SecretKey string
	ApiDomain string
	FeUrl     string
}

func NewAppConfig() *AppConfig {
	cnf := AppConfig{
		SecretKey: os.Getenv("SECRET"),
		ApiDomain: os.Getenv("API_DOMAIN"),
		FeUrl:     os.Getenv("FE_URL"),
	}
	return &cnf
}
