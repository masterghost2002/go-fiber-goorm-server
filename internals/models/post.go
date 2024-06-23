package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title     string
	Content   string
	UserId    int
	User      User
	CreatedAt time.Time
	UpdatedAt time.Time
}
