package category

type ExpenseCategory struct {
	ID          string
	DisplayName string
}

var OutsideFoodCategory = newOutsideFoodCategory()
var GrocerieCategory = newGroceriesCategory()
var CommutingCategory = newCommutingCategory()
var PurchasesCategory = newPurchasesCategory()
var UtilitiesCategory = newUtilitesCategory()
var MiscelleaneousCategory = newMiscellaneousCategory()
var ReimbursementCategory = newReimbursementCategory()

var categoryMap = map[string]ExpenseCategory{
	OutsideFoodCategory.ID:    *OutsideFoodCategory,
	GrocerieCategory.ID:       *GrocerieCategory,
	CommutingCategory.ID:      *CommutingCategory,
	PurchasesCategory.ID:      *PurchasesCategory,
	UtilitiesCategory.ID:      *UtilitiesCategory,
	MiscelleaneousCategory.ID: *MiscelleaneousCategory,
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

func newReimbursementCategory() *ExpenseCategory {
	return &ExpenseCategory{"reimbursement", "Reimbursement"}
}
