package db

import (
	"registry-auth/db/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Conn variable is singleton database connection
var Conn *gorm.DB

// Init function makes a database connection and migrates the models
func Init() {
	_db, err := gorm.Open("sqlite3", "test.db")
	Conn = _db

	if err != nil {
		panic("failed to connect database")
	}

	migrate()

	// defer db.Close()

}

func migrate() {
	Conn.AutoMigrate(&model.User{})
	Conn.AutoMigrate(&model.Repository{})
	Conn.AutoMigrate(&model.Tag{})
	Conn.AutoMigrate(&model.Access{})
	Conn.AutoMigrate(&model.AccessType{})
	Conn.AutoMigrate(&model.Action{})
}
