package models

import (
	"github.com/gargVader/telegram-expense-bot/app"
	"time"
)

type Expense struct {
	Date           time.Time `json:"date"`
	Description    string    `json:"description"`
	Amount         float64   `json:"amount"`
	CategoryID     string    `json:"categoryID"`
	SubCategoryID  string    `json:"subCategoryID"`
	IsReimbursable bool      `json:"isReimbursable"`
}

func (e *Expense) CreateExpense() {
	app.DB.Create(&e)
}

func GetAllExpenses() []Expense {
	var expenses []Expense
	app.DB.Find(expenses)
	return expenses
}

func GetExpenseById(ID int64) Expense {
	var expense Expense
	app.DB.Where("ID=?", ID).Find(&expense)
	return expense
}

func DeleteExpense(ID int64) {
	var e Expense
	app.DB.Where("ID=?", ID).Delete(e)
}

func (e *Expense) DeleteExpense() {
	app.DB.Delete(&e)
}
