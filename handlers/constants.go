package handlers

import (
	"fmt"
	"github.com/gargVader/telegram-expense-bot/models"
)

const wrong_spend_format_message = "😣 Wrong spend format. Correct format -> [Amount] [Description]"

func makeSuccessMessage(amount string, name string, categoryID string) string {
	category := models.CategoryMap[categoryID]
	return fmt.Sprintf("Expense Update! 📈\n\n✅ Successfully added <strong>¥%s</strong> spent at <strong>%s</strong> under <strong>%s</strong>", amount, name, category.DisplayName)
}

func makeWhichCategoryMessage(name string) string {
	return fmt.Sprintf("Which category does <strong>%s</strong> belong to?", name)
}
