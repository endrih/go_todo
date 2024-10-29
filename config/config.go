package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type DbSettings struct {
	DB_HOST     string
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

type GoogleAuthSettings struct {
	GOOGLE_OAUTH2_CLIENT_ID     string
	GOOGLE_OAUTH2_CLIENT_SECRET string
}

type AppConfig struct {
	SESSION_KEY  string
	GoogleConfig *GoogleAuthSettings
	IS_PROD      bool
	DbConfig     *DbSettings
}

func Initialize() *AppConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	isProd, err := strconv.ParseBool(os.Getenv("IS_PROD")) // Set to true when serving over https
	if err != nil {
		isProd = false
	}
	returnValue := &AppConfig{
		SESSION_KEY: os.Getenv("SESSION_KEY"),
		IS_PROD:     isProd,
		DbConfig: &DbSettings{
			DB_HOST:     os.Getenv("DB_HOST"),
			DB_USER:     os.Getenv("DB_USER"),
			DB_PASSWORD: os.Getenv("DB_PASSWORD"),
			DB_NAME:     os.Getenv("DB_NAME"),
		},
		GoogleConfig: &GoogleAuthSettings{
			GOOGLE_OAUTH2_CLIENT_ID:     os.Getenv("GOOGLE_OAUTH2_CLIENT_ID"),
			GOOGLE_OAUTH2_CLIENT_SECRET: os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET"),
		},
	}
	if returnValue.SESSION_KEY == "" {
		panic("SESSION_KEY not set")
	}

	if returnValue.DbConfig.DB_HOST == "" {
		panic("DB_HOST not set")
	}
	if returnValue.DbConfig.DB_NAME == "" {
		panic("DB_NAME not set")
	}
	if returnValue.DbConfig.DB_PASSWORD == "" {
		panic("DB_PASSWORD not set")
	}
	if returnValue.DbConfig.DB_USER == "" {
		panic("DB_USER not set")
	}
	if returnValue.GoogleConfig.GOOGLE_OAUTH2_CLIENT_ID == "" {
		panic("GOOGLE_OAUTH2_CLIENT_ID not set")
	}
	if returnValue.GoogleConfig.GOOGLE_OAUTH2_CLIENT_SECRET == "" {
		panic("GOOGLE_OAUTH2_CLIENT_SECRET not set")
	}
	return returnValue
}
