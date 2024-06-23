package db

import (
	"fmt"

	"github.com/masterghost2002/go-todo/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Storage *gorm.DB

func StorageInit() error {
	// dsn data source name
	dsn := "host=localhost user=root password=root dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
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
