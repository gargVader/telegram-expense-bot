package main

import (
	"log"
	"os"
	"time"

	"github.com/gargVader/telegram-expense-bot/handlers"
	"github.com/joho/godotenv"
	"gopkg.in/telebot.v3"
)

func init() {
	// Load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	pref := telebot.Settings{
		Token:  os.Getenv("TOKEN"),
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(telebot.OnText, handlers.OnTextHandler)
	b.Handle(telebot.OnCallback, handlers.CallbackHandler)

	b.Start()
}
