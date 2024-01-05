package models

import (
	"strings"
)

type DescriptionCategory struct {
	Description   string `json:"description" gorm:"primaryKey"`
	CategoryID    string `json:"categoryID"`
	SubCategoryID string `json:"subCategoryID"`
}

func (dc *DescriptionCategory) CreateDescriptionCategory() {
	GetDB().Create(&dc)
}

func GetAllDescriptionCategory() []DescriptionCategory {
	var d []DescriptionCategory
	GetDB().Find(d)
	return d
}

func GetDescriptionCategoryByDescription(description string) (DescriptionCategory, error) {
	description = strings.TrimSpace(description)
	description = strings.ToLower(description)
	var d DescriptionCategory
	result := GetDB().Where(&DescriptionCategory{Description: description}).First(&d)
	return d, result.Error
}

func DeleteDescriptionCategoryByDescription(description string) {
	var d DescriptionCategory
	GetDB().Where("description=?", description).Delete(d)
}

func (d *DescriptionCategory) DeleteDescriptionCategory() {
	GetDB().Delete(&d)
}
