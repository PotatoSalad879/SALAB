package config

import (
	"fmt"
	"os"

	"learnhub/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const defaultDSN = "host=localhost user=learnhub password=learnhub dbname=learnhub port=5432 sslmode=disable TimeZone=Asia/Bangkok"

// DSNFromEnv returns the database connection string from environment or the default DSN.
func DSNFromEnv() string {
	if dsn := os.Getenv("DATABASE_URL"); dsn != "" {
		return dsn
	}
	return defaultDSN
}

// Open opens a database connection with a custom DSN string.
func Open(databaseURL string) (*gorm.DB, error) {
	if databaseURL == "" {
		databaseURL = defaultDSN
	}

	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("open postgres connection: %w", err)
	}
	return db, nil
}

// AutoMigrate runs GORM auto-migrations for all LearnHub entities.
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Course{},
		&models.Tag{},
		&models.Enrollment{},
		&models.Lesson{},
		&models.Review{},
		&models.CourseSetting{},
	)
}

// ConnectDatabase opens a PostgreSQL connection and runs migrations using the provided Config.
func ConnectDatabase(cfg *Config) (*gorm.DB, error) {
	dsn := DSNFromEnv()
	if os.Getenv("DATABASE_URL") == "" && cfg != nil {
		dsn = cfg.DSN()
	}

	db, err := Open(dsn)
	if err != nil {
		return nil, err
	}

	if err := AutoMigrate(db); err != nil {
		return nil, fmt.Errorf("database auto-migration failed: %w", err)
	}
	return db, nil
}
