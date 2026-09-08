package db

import (
	"fmt"
	"os"

	"skylink/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const defaultDSN = "host=localhost user=skylink password=skylink dbname=skylink port=5432 sslmode=disable TimeZone=Asia/Bangkok"

// DSN returns the connection string from DATABASE_URL or the default.
func DSN() string {
	if v := os.Getenv("DATABASE_URL"); v != "" {
		return v
	}
	return defaultDSN
}

// Open connects to PostgreSQL.
func Open(dsn string) (*gorm.DB, error) {
	if dsn == "" {
		dsn = defaultDSN
	}
	gdb, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, fmt.Errorf("open postgres: %w", err)
	}
	return gdb, nil
}

// AutoMigrate builds the full schema from the model structs.
func AutoMigrate(gdb *gorm.DB) error {
	return gdb.AutoMigrate(
		&models.Airport{},
		&models.Aircraft{},
		&models.Seat{},
		&models.Flight{},
		&models.CrewMember{},
		&models.FlightCrew{},
		&models.Passenger{},
		&models.Note{},
		&models.Booking{},
		&models.Ticket{},
	)
}
