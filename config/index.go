package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	AppPort     string
	DbPath      string
	AutoMigrate bool
	JwtSecret   string
}

var Config envConfig

func (e *envConfig) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}

	portValue, ok := os.LookupEnv("APP_PORT")
	if !ok {
		log.Panic("APP_PORT not set in environment")
	}

	dbPathValue, ok := os.LookupEnv("DB_PATH")
	if !ok {
		log.Panic("DB_PATH not set in environment")
	}

	jwtSecretValue, ok := os.LookupEnv("JWT_SECRET")
	if !ok {
		log.Panic("JWT_SECRET not set in environment")
	}

	e.AppPort = portValue
	e.DbPath = dbPathValue
	e.JwtSecret = jwtSecretValue

	autoMigrateValue, ok := os.LookupEnv("AUTO_MIGRATE")
	if !ok {
		log.Panic("AUTO_MIGRATE not set in environment")
	}

	e.AutoMigrate = autoMigrateValue == "true"
}

func init() {
	Config.LoadConfig()
}
