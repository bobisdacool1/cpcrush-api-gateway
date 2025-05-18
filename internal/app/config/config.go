package config

import (
	"os"
)

type (
	AppConfig struct {
		DBConn string
	}
)

func NewConfig() AppConfig {
	return AppConfig{
		DBConn: os.Getenv("DB_CONN"),
	}
}
