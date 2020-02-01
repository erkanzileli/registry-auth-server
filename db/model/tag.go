package model

import "github.com/jinzhu/gorm"

// Tag model
type Tag struct {
	gorm.Model
	Name       string
	Repository Repository
}
