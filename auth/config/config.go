package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbMigrationDir     string `envconfig:"DB_MIGRATION_DIR"`
	DbHost             string `envconfig:"DB_HOST"`
	DbPort             int    `envconfig:"DB_PORT"`
	DBName             string `envconfig:"DB_NAME"`
	DbUser             string `envconfig:"DB_USER"`
	DbPassword         string `envconfig:"DB_PASSWORD"`
	JWTSigningKey      string `envconfig:"JWT_SIGNING_KEY"`
	CaesarCipherOffset int    `envconfig:"CAESAR_CIPHER_OFFSET"`
}

func FromEnv() Config {
	err := autoLoadEnv()
	if err != nil {
		panic(err)
	}

	config := Config{}
	err = envconfig.Process("", &config)
	if err != nil {
		panic(err)
	}
	return config
}

func autoLoadEnv() error {
	_, err := os.Stat(".env")
	if os.IsNotExist(err) {
		return nil
	}

	return godotenv.Load()
}
