package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	DB *gorm.DB
}

func NewConfig() (*Config, error) {
	db, err := setupDatabase()
	if err != nil {
		return nil, fmt.Errorf("failed to setup database: %w", err)
	}

	return &Config{
		DB: db,
	}, nil
}

func setupDatabase() (*gorm.DB, error) {
	// Get database configuration from environment variables with defaults
	host := getEnv("DB_HOST", "localhost")
	port := getEnv("DB_PORT", "5432")
	user := getEnv("DB_USER", "postgres")
	password := getEnv("DB_PASSWORD", "postgres")
	dbname := getEnv("DB_NAME", "cat_feeder")

	// Retry connection logic
	var db *gorm.DB
	var err error
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			host, user, password, dbname, port,
		)

		db, err = connectWithConfig(dsn)
		if err == nil {
			break
		}

		log.Printf("Failed to connect to database (attempt %d/%d): %v", i+1, maxRetries, err)
		time.Sleep(time.Second * 5)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database after %d attempts: %w", maxRetries, err)
	}

	return db, nil
}

func connectWithConfig(dsn string) (*gorm.DB, error) {
	// Custom logger configuration
	logConfig := logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logConfig,
		),
	})
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
