package handlers

import (
	"github.com/gargVader/telegram-expense-bot/log"
	"github.com/gargVader/telegram-expense-bot/models"
	"gopkg.in/telebot.v3"
	"strconv"
	"strings"
	"time"
)

func OnTextHandler(c telebot.Context) error {
	text := c.Text()
	log.Logger.Printf("OnTextHandler: %s", text)

	return addSpend(c)
}

func addSpend(c telebot.Context) error {
	var (
		text = c.Text()
	)

	// Number of words is less than 2
	vals := strings.Split(text, " ")
	if len(vals) < 2 {
		return c.Send(wrong_spend_format_message)
	}

	// First word is not a Float
	amount, err := strconv.ParseFloat(vals[0], 64)
	if err != nil {
		return c.Send(wrong_spend_format_message)
	}

	description := strings.Join(vals[1:], " ")

	log.Logger.Printf("Amount=%g, Name=%s\n", amount, description)

	// Form Response
	dc, err := models.GetDescriptionCategoryByDescription(description)
	if err == nil {
		// Category already known
		// Save to database
		newExpense := models.Expense{
			Date:          time.Now(),
			Description:   description,
			Amount:        amount,
			CategoryID:    dc.CategoryID,
			SubCategoryID: dc.SubCategoryID,
		}
		newExpense.CreateExpense()
		resp := makeSuccessMessage(strconv.FormatFloat(amount, 'f', -1, 64), description, dc.CategoryID)
		return c.EditOrSend(resp, "HTML")
	} else {
		// Category not known
		resp := makeWhichCategoryMessage(description)
		selector := &telebot.ReplyMarkup{}
		rowList := mapExpenseCategoryToSelectorRowList(*selector, models.ExpenseCategories, amount, strings.TrimSpace(description))
		selector.Inline(
			rowList...,
		)
		return c.EditOrSend(resp, selector, "HTML")
	}
}

func mapExpenseCategoryToSelectorRowList(selector telebot.ReplyMarkup, expenseCategoryList []models.Category, amount float64, name string) []telebot.Row {
	var rowList []telebot.Row
	for _, category := range expenseCategoryList {
		row := selector.Row(selector.Data(category.DisplayName, category.ID, strconv.FormatFloat(amount, 'f', -1, 64), name))
		rowList = append(rowList, row)
	}
	return rowList
}
