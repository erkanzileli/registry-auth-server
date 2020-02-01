package model

import "github.com/jinzhu/gorm"

// Action model
type Action struct {
	gorm.Model
	Name string
}
