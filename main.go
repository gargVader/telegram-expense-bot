package main

import (
	"github.com/gargVader/telegram-expense-bot/models"
	"os"
	"time"

	"github.com/gargVader/telegram-expense-bot/handlers"
	"github.com/gargVader/telegram-expense-bot/log"
	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func init() {
	// Init logger
	log.Init()
	// Init DB
	models.ConnectDB()
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Logger.Fatal("Error loading .env file")
	}
}

func main() {
	// Setup telegram bot
	pref := telebot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(pref)
	if err != nil {
		log.Logger.Fatal(err)
		return
	}

	bot.Handle(telebot.OnText, handlers.OnTextHandler)
	bot.Handle(telebot.OnCallback, handlers.CallbackHandler)

	bot.Start()
}
