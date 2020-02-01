package model

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string
}
