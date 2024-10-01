package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Env string `env:"ENV" envDefault:"local"`
	HTTPServer
	PgDb
	Migrations
}

type HTTPServer struct {
	Address string `env:"HTTP_ADDRESS" envDefault:"8080"`
}

type PgDb struct {
	PGName     string `env:"PG_DATABASE" envDefault:"effectiveMode_db"`
	PGPassword string `env:"PG_PASSWORD" envDefault:"passwordd"`
	PGUser     string `env:"PG_USER" envDefault:"newuser"`
	PGHost     string `env:"PG_HOST" envDefault:"localhost"`
	PGPort     string `env:"PG_PORT" envDefault:"5432"`
}

type Migrations struct {
	MigrationPath string
}

func MustLoad() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	return &Config{
		Env: os.Getenv("ENV"),
		HTTPServer: HTTPServer{
			Address: os.Getenv("HTTP_ADDRESS"),
		},
		PgDb: PgDb{
			PGName:     os.Getenv("PG_DATABASE"),
			PGPassword: os.Getenv("PG_PASSWORD"),
			PGUser:     os.Getenv("PG_USER"),
			PGHost:     os.Getenv("PG_HOST"),
			PGPort:     os.Getenv("PG_PORT"),
		},
	}
}
