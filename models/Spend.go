package models

import (
	"time"
)

type Expense struct {
	Date          time.Time `json:"date"`
	Description   string    `json:"description"`
	Amount        float64   `json:"amount"`
	CategoryID    string    `json:"categoryId"`
	SubCategoryID string    `json:"subCategoryId"`
}

func (e *Expense) CreateExpense() {
	GetDB().Create(&e)
}

func GetAllExpenses() []Expense {
	var expenses []Expense
	GetDB().Find(expenses)
	return expenses
}

func GetExpenseById(ID int64) Expense {
	var expense Expense
	GetDB().Where("ID=?", ID).Find(&expense)
	return expense
}

func DeleteExpense(ID int64) {
	var e Expense
	GetDB().Where("ID=?", ID).Delete(e)
}

func (e *Expense) DeleteExpense() {
	GetDB().Delete(&e)
}
