package category

type ExpenseCategory struct {
	Name        string
	DisplayName string
}

var OutsideFoodCategory = newOutsideFoodCategory()
var GrocerieCategory = newGroceriesCategory()
var CommutingCategory = newCommutingCategory()
var PurchasesCategory = newPurchasesCategory()
var UtilitiesCategory = newUtilitesCategory()
var MiscelleaneousCategory = newMiscellaneousCategory()

var categoryMap = map[string]ExpenseCategory{
	OutsideFoodCategory.Name:    *OutsideFoodCategory,
	GrocerieCategory.Name:       *GrocerieCategory,
	CommutingCategory.Name:      *CommutingCategory,
	PurchasesCategory.Name:      *PurchasesCategory,
	UtilitiesCategory.Name:      *UtilitiesCategory,
	MiscelleaneousCategory.Name: *MiscelleaneousCategory,
}

var ExpenseCategories = []ExpenseCategory{
	*OutsideFoodCategory,
	*GrocerieCategory,
	*CommutingCategory,
	*PurchasesCategory,
	*UtilitiesCategory,
	*MiscelleaneousCategory,
}

func newOutsideFoodCategory() *ExpenseCategory {
	return &ExpenseCategory{"outside_food", "🍔 Outside Food"}
}

func newGroceriesCategory() *ExpenseCategory {
	return &ExpenseCategory{"groceries", "🛒 Groceries"}
}

func newCommutingCategory() *ExpenseCategory {
	return &ExpenseCategory{"commuting", "🚉 Commuting"}
}

func newPurchasesCategory() *ExpenseCategory {
	return &ExpenseCategory{"purchases", "👕👟🎧 Purchases"}
}

func newUtilitesCategory() *ExpenseCategory {
	return &ExpenseCategory{"utilities", "📱🛜🔌🚰⛽️Mobile, Utilites"}
}

func newMiscellaneousCategory() *ExpenseCategory {
	return &ExpenseCategory{"miscelleaneous", "Miscelleaneous"}
}
