package main

import (
	"github.com/gargVader/telegram-expense-bot/app"
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
	if err := app.ConnectDB(&models.Expense{}, &models.DescriptionCategory{}); err != nil {
		log.Logger.Fatal(err)
		return
	}
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Logger.Fatal("Error loading .env file")
		return
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
	bot.Handle(telebot.OnCallback, handlers.OnCallbackHandler)

	bot.Start()
}

/*
TODOS
- Refine logs
- Add command to delete last entry
- Identify any potential error handling
- Study options given by Telegram

Features:
- /delete
- /stats
- Add records retroactively. Yesterday ...
- Add rows to Google Sheets
- Add subcategories
- Figure out solution for reimbursement

Deployment:
- Dockerize
- Connect with an SQL dashboard. [This should be done only after importing old data]
*/
