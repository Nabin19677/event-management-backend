package models

type EventExpenseFields struct {
	EventID     int     `json:"eventId" db:"event_id"`
	ItemName    string  `json:"itemName" db:"item_name"`
	Cost        float64 `json:"cost" db:"cost"`
	Description string  `json:"description" db:"description"`
	CategoryID  int     `json:"categoryId" db:"category_id"`
}

type EventExpense struct {
	ExpenseID int `json:"expenseId" db:"expense_id"`
	EventExpenseFields
}

type NewEventExpense struct {
	EventExpenseFields
}

type EventExpenseCategory struct {
	CategoryID   int    `json:"categoryId" db:"category_id"`
	CategoryName string `json:"categoryName" db:"category_name"`
}

type CategoryTotal struct {
	CategoryName string  `json:"categoryName" db:"category_name"`
	TotalCost    float64 `json:"totalCost"  db:"total_cost"`
}
