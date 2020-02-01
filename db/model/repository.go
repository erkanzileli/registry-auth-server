package model

import "github.com/jinzhu/gorm"

// Repository model
type Repository struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Tags []Tag
}
