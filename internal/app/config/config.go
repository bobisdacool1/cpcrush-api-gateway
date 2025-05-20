package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type (
	AppConfig struct {
		DBConn string `yaml:"db_conn"`
		Port   string `yaml:"port"`
	}
)

func MustNewConfig() AppConfig {
	cfg, err := newConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg
}

func newConfig() (AppConfig, error) {
	path := getConfigPath()
	file, err := os.ReadFile(path) //nolint:gosec
	if err != nil {
		return AppConfig{}, errors.Wrap(err, "read config file")
	}

	var cfg AppConfig
	if err = yaml.Unmarshal(file, &cfg); err != nil {
		return AppConfig{}, errors.Wrap(err, "unmarshal config")
	}

	return cfg, nil
}

func getConfigPath() string {
	if path := os.Getenv("CONFIG_PATH"); path != "" {
		return path
	}
	return filepath.Join("configs", "config.yaml")
}
