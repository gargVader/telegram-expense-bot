package handlers

import (
	"fmt"
	"strconv"
	"strings"

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

	vals := strings.Split(text, " ")
	if len(vals) < 2 {
		return c.Send("Wrong spend format")
	}

	amount, err := strconv.ParseFloat(vals[0], 64)
	if err != nil {
		return c.Send("Wrong spend format")
	}

	name := strings.Join(vals[1:], " ")

	fmt.Printf("Amount=%g, Name=%s\n", amount, name)

	// Form Response
	resp := fmt.Sprintf("Which category does <strong>%s</strong> belong to?", name)
	selector := &telebot.ReplyMarkup{}
	selector.Inline(
		selector.Row(selector.Data("🍔 Outside Food", "outside_food", strconv.FormatFloat(amount, 'f', -1, 64), name)),
		selector.Row(selector.Data("🛒 Groceries", "groceries")),
		selector.Row(selector.Data("🚉 Commuting", "commuting")),
		selector.Row(selector.Data("👕👟🎧Clothes & Electronics", "utilities")),
		selector.Row(selector.Data("📱🛜🔌🚰⛽️Mobile, Utilites", "utilities")),
		selector.Row(selector.Data("Miscelleaneous", "miscelleaneous")),
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
	resp := fmt.Sprintf("Expense Update! 📈\n\n✅ Successfully added <strong>¥%s</strong> spent at <strong>%s</strong> under %s", amount, name, category)
	return c.EditOrSend(resp, "HTML")

}
