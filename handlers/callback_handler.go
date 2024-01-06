package handlers

import (
	"github.com/gargVader/telegram-expense-bot/log"
	"github.com/gargVader/telegram-expense-bot/models"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
	"time"
)

func OnCallbackHandler(c telebot.Context) error {
	var (
		args = c.Args()
		data = c.Data()
	)
	log.Logger.Printf("OnCallbackHandler: %v", args)

	category := strings.TrimSpace(args[0])
	amount := args[1]
	description := args[2]

	amountFloat, _ := strconv.ParseFloat(amount, 64)

	// Save to database
	newExpense := models.Expense{
		Date:        time.Now(),
		Description: description,
		Amount:      amountFloat,
		CategoryID:  category,
	}
	newExpense.CreateExpense()
	resp := makeSuccessMessage(amount, description, category)
	// Also add to DescriptionCategory table
	description = strings.TrimSpace(description)
	description = strings.ToLower(description)
	newDC := models.DescriptionCategory{
		Description: description,
		CategoryID:  category,
	}
	newDC.CreateDescriptionCategory()
	return c.EditOrSend(resp, "HTML")
}
