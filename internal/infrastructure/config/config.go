package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	PORT string `split_words:"PORT" default:"4000"`

	DB_URL         string `split_words:"DB_URL"`
	DB_HOST        string `split_words:"DB_HOST"`
	DB_PORT        int    `split_words:"DB_PORT"`
	DB_USER        string `split_words:"DB_USER"`
	DB_PASSWORD    string `split_words:"DB_PASSWORD"`
	DB_NAME        string `split_words:"DB_NAME"`
	DB_SSLMODE     string `split_words:"DB_SSLMODE"`
	DB_SCHEMA_NAME string `split_words:"DB_SCHEMA_NAME"`

	// Migrations
	PathMigrations  string `split_words:"true" default:"migrations"`
	MigrationDBAuto bool   `split_words:"true" default:"false"`
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, fmt.Errorf("failed to load .env file: %w", err)
	}

	var cfg Config
	if err := envconfig.Process("", &cfg); err != nil {
		return nil, fmt.Errorf("failed to process config: %w", err)
	}

	return &cfg, nil
}
