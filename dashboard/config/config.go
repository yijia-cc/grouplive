package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbDriver     string `envconfig:"DB_DRIVER"`
	DbHost         string `envconfig:"DB_HOST"`
	DbPort         string `envconfig:"DB_PORT"`
	DBName         string `envconfig:"DB_NAME"`
	DbUser         string `envconfig:"DB_USER"`
	DbPassword     string `envconfig:"DB_PASSWORD"`
	TokenSecretKey string `envconfig:"TOKEN_SECRET_KEY"`
}


func LoadEnv() (*Config, *Config) {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	userDBConfig := Config{}
	if err := envconfig.Process("USER", &userDBConfig); err != nil {
		panic(err)
	}

	dashDBConfig := Config{}
	if err := envconfig.Process("DASH", &dashDBConfig); err != nil {
		panic(err)
	}

	//fmt.Println("userDBConfig: ", userDBConfig)
	//fmt.Println("dashDBConfig: ", dashDBConfig)

	return &userDBConfig, &dashDBConfig
}
