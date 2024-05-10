package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	App  *App
	Db   *Db
	Grpc *Grpc
}

type App struct {
	Url     string
	AppName string
}

type Db struct {
	Url string
}

type Grpc struct {
	TaskAppUrl string
}

func NewConfig(path string) *Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		App: &App{
			Url:     os.Getenv("APP_URL"),
			AppName: os.Getenv("APP_NAME"),
		},
		Db: &Db{
			Url: os.Getenv("DB_URL"),
		},
		Grpc: &Grpc{
			TaskAppUrl: os.Getenv("GRPC_TASK_APP_URL"),
		},
	}
}
