package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server   serverConfig
	Database databaseConfig
	Api      apiConfig
}

type databaseConfig struct {
	MongoUrl       string
	NameDb         string
	NameCollection string
}

type serverConfig struct {
	Host string
	Port string
}

type apiConfig struct {
	Url string
}

func NewConfig() Config {
	if err := godotenv.Load(".env"); err != nil {
		return Config{}
	}

	return Config{
		serverConfig{
			Host: os.Getenv("SERVER_HOST"),
			Port: os.Getenv("SERVER_PORT"),
		},
		databaseConfig{
			MongoUrl:       os.Getenv("MONGODB_URL"),
			NameDb:         os.Getenv("NAME_DB"),
			NameCollection: os.Getenv("NAME_COLLECTION"),
		},
		apiConfig{},
	}
}
