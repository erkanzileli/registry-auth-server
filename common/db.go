package common

import (
	"registry-auth/db/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB variable is singleton database connection
var DB *gorm.DB

// Init function makes a database connection and migrates the models
func Init() {
	_db, err := gorm.Open("sqlite3", "test.db")
	DB = _db

	if err != nil {
		panic("failed to connect database")
	}

	migrate()

	// defer db.Close()

}

func migrate() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Repository{})
	DB.AutoMigrate(&model.Tag{})
	DB.AutoMigrate(&model.Access{})
	DB.AutoMigrate(&model.AccessType{})
	DB.AutoMigrate(&model.Action{})
}
