package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/BAITOEYSRN/test-Technical-Skill/internal/infrastructure/config"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MigrationConfig struct {
	Path string
}

func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	rawDB, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
			cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_SSLMODE))
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}
	defer rawDB.Close()

	_, err = rawDB.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", cfg.DB_SCHEMA_NAME))
	if err != nil {
		return nil, fmt.Errorf("failed to create schema: %w", err)
	}

	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s search_path=%s",
		cfg.DB_HOST, cfg.DB_PORT, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_SSLMODE, cfg.DB_SCHEMA_NAME,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect database with GORM: %w", err)
	}

	log.Println("Connected to database successfully!")
	return db, nil
}

func MigrateDB(cfg *config.Config, migrations []MigrationConfig) error {
	for _, migration := range migrations {
		databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s&search_path=%s&x-migrations-table=schema_migrations",
			cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME, cfg.DB_SSLMODE, cfg.DB_SCHEMA_NAME,
		)
		m, err := migrate.New(
			fmt.Sprintf("file://%s", migration.Path),
			databaseURL,
		)
		if err != nil {
			return fmt.Errorf("failed to migrate database: %w", err)
		}
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			return fmt.Errorf("failed to run migrations: %w", err)
		}
	}
	return nil
}
