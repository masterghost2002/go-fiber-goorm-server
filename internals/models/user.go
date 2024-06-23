package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  *string // get auto converted to last_name
	Email     string  `gorm:"unique;not null"`
	Password  string  `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
