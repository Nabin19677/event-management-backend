package models

type EventExpense struct {
	ExpenseID   int     `json:"expenseId" db:"expense_id"`
	EventID     int     `json:"eventId" db:"event_id"`
	ItemName    string  `json:"itemName" db:"item_name"`
	Cost        float64 `json:"cost" db:"cost"`
	Description string  `json:"description" db:"description"`
	CategoryID  int     `json:"categoryId" db:"category_id"`
}

type EventExpenseCategory struct {
	CategoryID   int    `json:"categoryId" db:"category_id"`
	CategoryName string `json:"categoryName" db:"category_name"`
}
