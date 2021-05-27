package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"os"
)

type Config struct {
	DbDriver       string `envconfig:"DB_DRIVER"`
	DbHost         string `envconfig:"DB_HOST"`
	DbPort         string `envconfig:"DB_PORT"`
	DBName         string `envconfig:"DB_NAME"`
	DbUser         string `envconfig:"DB_USER"`
	DbPassword     string `envconfig:"DB_PASSWORD"`
	TokenSecretKey string `envconfig:"TOKEN_SECRET_KEY"`
}


func LoadEnv() *Config {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		panic(err)
	}

	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	config := &Config{}
	if err := envconfig.Process("", config); err != nil {
		panic(err)
	}
	return config
}