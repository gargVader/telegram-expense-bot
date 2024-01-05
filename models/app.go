package models

import (
	"github.com/gargVader/telegram-expense-bot/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	log.Logger.Print("Connecting to db...")
	d, err := gorm.Open(sqlite.Open("./data/data.db"), &gorm.Config{})
	if err != nil {
		log.Logger.Fatal("Failed to connect database")
	}
	db = d
	db.AutoMigrate(&Expense{})
	db.AutoMigrate(&DescriptionCategory{})
}

func GetDB() *gorm.DB {
	return db
}
