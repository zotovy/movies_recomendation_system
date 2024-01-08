package db

import (
	"app/pkg/config"
	"app/pkg/logger"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

type DB = gorm.DB

func NewDatabase(config *config.Config, logger *logger.Logger) *DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		config.PostgreSQL.Host,
		config.PostgreSQL.Username,
		config.PostgreSQL.Password,
		config.PostgreSQL.Database,
		config.PostgreSQL.Port,
	)

	dbLogger := zapgorm2.New(logger)

	logger.Info("Connecting to PostgreSQL", zap.String("dsn", dsn))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: dbLogger})

	if err != nil {
		logger.Fatal("Failed to connect to PostgreSQL", zap.Error(err))
		return nil
	}

	// Migrations
	//err = db.AutoMigrate(&models.Test{})

	if err != nil {
		logger.Fatal("Failed to auto migrate database", zap.Error(err))
		return nil
	}

	return db
}
