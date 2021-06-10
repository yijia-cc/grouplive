package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

var Cfg *Config

type Config struct {
	App *appConfig
	Db  map[string]*DbConfig
}

type DbConfig struct {
	DbDriver      string `envconfig:"DB_DRIVER"`
	DbHost        string `envconfig:"DB_HOST"`
	DbPort        int    `envconfig:"DB_PORT"`
	DBName        string `envconfig:"DB_NAME"`
	DbUser        string `envconfig:"DB_USER"`
	DbPassword    string `envconfig:"DB_PASSWORD"`
	JwtSigningKey string `envconfig:"JWT_SIGNING_KEY"`
}

type appConfig struct {
	WebServerHost       string `envconfig:"APP_WEB_SERVER_HOST"`
	WebServerPort       int    `envconfig:"APP_WEB_SERVER_PORT"`
	JwtSigningKey       string `envconfig:"APP_JWT_SIGNING_KEY"`
	StaticMediaDir      string `envconfig:"APP_STATIC_MEDIA_DIR"`
	FileUploadSizeLimit int64  `envconfig:"APP_FILE_UPLOAD_SIZE_LIMIT"`
	LocalDatetimeFormat string `envconfig:"APP_LOCAL_DATETIME_FORMAT"`
}

func FromEnv() *Config {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	userDbCfg := DbConfig{}
	if err := envconfig.Process("USER", &userDbCfg); err != nil {
		panic(err)
	}

	dashDbCfg := DbConfig{}
	if err := envconfig.Process("DASHBOARD", &dashDbCfg); err != nil {
		panic(err)
	}

	appCfg := appConfig{}
	if err := envconfig.Process("APP", &appCfg); err != nil {
		panic(err)
	}

	//fmt.Println("userDbCfg:", userDbCfg)
	//fmt.Println("dashDbCfg:", dashDbCfg)

	Cfg = &Config{
		App: &appCfg,
		Db: map[string]*DbConfig{
			"user": &userDbCfg,
			"dash": &dashDbCfg,
		},
	}
	return Cfg
}