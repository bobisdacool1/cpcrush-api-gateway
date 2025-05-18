package config

import (
	"os"
)

type (
	AppConfig struct {
		DbConn string
	}
)

func NewConfig() AppConfig {
	return AppConfig{
		DbConn: os.Getenv("DB_CONN"),
	}
}
