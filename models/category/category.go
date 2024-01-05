package category

type Category struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type SubCategory struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	CategoryID  string `json:"categoryId"`
}

var OutsideFoodCategory = newOutsideFoodCategory()
var GroceriesCategory = newGroceriesCategory()
var CommutingCategory = newCommutingCategory()
var PurchasesCategory = newPurchasesCategory()
var UtilitiesCategory = newUtilitesCategory()
var MiscellaneousCategory = newMiscellaneousCategory()
var ReimbursementCategory = newReimbursementCategory()

var categoryMap = map[string]Category{
	OutsideFoodCategory.ID:   *OutsideFoodCategory,
	GroceriesCategory.ID:     *GroceriesCategory,
	CommutingCategory.ID:     *CommutingCategory,
	PurchasesCategory.ID:     *PurchasesCategory,
	UtilitiesCategory.ID:     *UtilitiesCategory,
	MiscellaneousCategory.ID: *MiscellaneousCategory,
	ReimbursementCategory.ID: *ReimbursementCategory,
}

var ExpenseCategories = []Category{
	*OutsideFoodCategory,
	*GroceriesCategory,
	*CommutingCategory,
	*PurchasesCategory,
	*UtilitiesCategory,
	*MiscellaneousCategory,
	*ReimbursementCategory,
}

func newOutsideFoodCategory() *Category {
	return &Category{"outside_food", "🍔 Outside Food"}
}

func newGroceriesCategory() *Category {
	return &Category{"groceries", "🛒 Groceries"}
}

func newCommutingCategory() *Category {
	return &Category{"commuting", "🚉 Commuting"}
}

func newPurchasesCategory() *Category {
	return &Category{"purchases", "👕👟🎧 Purchases"}
}

func newUtilitesCategory() *Category {
	return &Category{"utilities", "📱🛜🔌🚰⛽️Mobile, Utilites"}
}

func newMiscellaneousCategory() *Category {
	return &Category{"miscelleaneous", "Miscelleaneous"}
}

func newReimbursementCategory() *Category {
	return &Category{"reimbursement", "Reimbursement"}
}
