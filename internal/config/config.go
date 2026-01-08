package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBDNS string
	Port  string
}

func Load() *Config {
	return &Config{
		DBDNS: fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		),
		Port: os.Getenv("PORT"),
	}
}
