package model

import "github.com/jinzhu/gorm"

// Access model
type Access struct {
	gorm.Model
	Name    string
	Type    AccessType
	Actions []Action
}
