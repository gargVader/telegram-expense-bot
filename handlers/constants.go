package handlers

import (
	"fmt"
)

const wrong_spend_format_message = "😣 Wrong spend format. Correct format -> [Amount] [Description]"

func makeSuccessMessage(amount string, name string, category string) string {
	return fmt.Sprintf("Expense Update! 📈\n\n✅ Successfully added <strong>¥%s</strong> spent at <strong>%s</strong> under %s", amount, name, category)
}

func makeWhichCategoryMessage(name string) string {
	return fmt.Sprintf("Which category does <strong>%s</strong> belong to?", name)
}
