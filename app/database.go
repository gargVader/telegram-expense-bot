package app

import (
	"github.com/gargVader/telegram-expense-bot/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dst ...interface{}) error {
	log.Logger.Print("Connecting to databse...")
	database, err := gorm.Open(sqlite.Open("./data/data.DB"), &gorm.Config{})
	if err != nil {
		log.Logger.Fatal("Failed to connect database")
		return err
	}
	DB = database
	log.Logger.Print("Connected to database")
	return DB.AutoMigrate(dst...)
}
