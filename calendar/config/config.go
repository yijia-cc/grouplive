package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DbHost         string `envconfig:"DB_HOST"`
	DbPort         int    `envconfig:"DB_PORT"`
	DbUser         string `envconfig:"DB_USER"`
	DbPassword     string `envconfig:"DB_PASSWORD"`
	DBName         string `envconfig:"DB_NAME"`
	DbMigrationDir string `envconfig:"DB_MIGRATION_DIR"`
	GraphQLSchemaPath string `envconfig:"GRAPHQL_SCHEMA_PATH"`
	GraphQLServerPort int `envconfig:"GRAPHQL_SERVER_PORT"`
}

func FromEnv() Config{
	err := godotenv.Load()
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