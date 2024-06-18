package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbHost string
	DbPort string
	DbUser string
	DbPass string
	DbName string
}

func LoadConfig() Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading config file")
	}

	return Config{
		DbHost: os.Getenv("DB_HOST"),
		DbPort: os.Getenv("DB_PORT"),
		DbUser: os.Getenv("DB_USER"),
		DbPass: os.Getenv("DB_PASS"),
		DbName: os.Getenv("DB_NAME"),
	}
}
