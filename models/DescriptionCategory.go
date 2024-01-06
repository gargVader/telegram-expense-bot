package models

import (
	"github.com/gargVader/telegram-expense-bot/app"
	"strings"
)

type DescriptionCategory struct {
	Description   string `json:"description" gorm:"primaryKey"`
	CategoryID    string `json:"categoryID"`
	SubCategoryID string `json:"subCategoryID"`
}

func (dc *DescriptionCategory) CreateDescriptionCategory() {
	app.DB.Create(&dc)
}

func GetAllDescriptionCategory() []DescriptionCategory {
	var d []DescriptionCategory
	app.DB.Find(d)
	return d
}

func GetDescriptionCategoryByDescription(description string) (DescriptionCategory, error) {
	description = strings.TrimSpace(description)
	description = strings.ToLower(description)
	var d DescriptionCategory
	result := app.DB.Where(&DescriptionCategory{Description: description}).First(&d)
	return d, result.Error
}

func DeleteDescriptionCategoryByDescription(description string) {
	var d DescriptionCategory
	app.DB.Where("description=?", description).Delete(d)
}

func (d *DescriptionCategory) DeleteDescriptionCategory() {
	app.DB.Delete(&d)
}
