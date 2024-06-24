package db

import (
	"fmt"

	config "github.com/masterghost2002/go-todo/configs"
	"github.com/masterghost2002/go-todo/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Storage *gorm.DB

func StorageInit() error {
	// dsn data source name
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.ENVS.DBAddress, config.ENVS.DBUser, config.ENVS.DBPassword, config.ENVS.DBName, config.ENVS.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	fmt.Println("Successfully connected to database")
	if err := db.AutoMigrate(&models.User{}, &models.Post{}); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	Storage = db
	return nil
}
