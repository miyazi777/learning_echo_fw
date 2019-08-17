package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func getDbConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "/tmp/test.db")
	if err != nil {
		panic("Failed to connect DB.")
	}
	return db
}

func InitDb() {
	db := getDbConnection()
	defer db.Close()

	db.AutoMigrate(&Item{})
}
