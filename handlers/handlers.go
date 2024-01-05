package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gargVader/telegram-expense-bot/models/category"
	"gopkg.in/telebot.v3"
)

func OnTextHandler(c telebot.Context) error {
	text := c.Text()
	fmt.Println(text)

	return AddSpendHandler(c)
}

func getSelectorRows() {

}

func AddSpendHandler(c telebot.Context) error {
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

	name := strings.Join(vals[1:], " ")

	fmt.Printf("Amount=%g, Name=%s\n", amount, name)

	// Form Response
	resp := makeWhichCategoryMessage(name)
	selector := &telebot.ReplyMarkup{}
	rowList := mapExpenseCategoryToSelectorRowList(*selector, category.ExpenseCategories, amount, name)
	selector.Inline(
		rowList...,
	)

	return c.EditOrSend(resp, selector, "HTML")
}

func CallbackHandler(c telebot.Context) error {
	var (
		args = c.Args()
	)

	category := args[0]
	amount := args[1]
	name := args[2]

	fmt.Println(args[0])

	// Save to database

	resp := makeSuccessMessage(amount, name, category)
	return c.EditOrSend(resp, "HTML")
}

func mapExpenseCategoryToSelectorRowList(selector telebot.ReplyMarkup, expenseCategoryList []category.ExpenseCategory, amount float64, name string) []telebot.Row {
	var rowList []telebot.Row
	for _, category := range expenseCategoryList {
		row := selector.Row(selector.Data(category.DisplayName, category.Name, strconv.FormatFloat(amount, 'f', -1, 64), name))
		rowList = append(rowList, row)
	}
	return rowList
}
