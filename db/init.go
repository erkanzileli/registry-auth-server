package db

import (
	"registry-auth/db/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// Init function makes a database connection and migrates the models
func Init() {
	_db, err := gorm.Open("sqlite3", "test.db")
	db = _db

	if err != nil {
		panic("failed to connect database")
	}

	migrate()

	// defer db.Close()

}

func migrate() {
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Repository{})
	db.AutoMigrate(&model.Tag{})
	db.AutoMigrate(&model.Access{})
	db.AutoMigrate(&model.AccessType{})
	db.AutoMigrate(&model.Action{})
}
