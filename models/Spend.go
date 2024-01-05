package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Expense struct {
	gorm.Model
	Date        time.Time
	Description string
	Amount      float64
	category    ExpenseCategory
}
