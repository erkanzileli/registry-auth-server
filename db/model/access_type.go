package model

import "github.com/jinzhu/gorm"

// AccessType model
type AccessType struct {
	gorm.Model
	Name string
}
