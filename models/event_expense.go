package models

type EventExpenseFields struct {
	EventID     int     `json:"eventId" db:"event_id" validate:"required"`
	ItemName    string  `json:"itemName" db:"item_name" validate:"required"`
	Cost        float64 `json:"cost" db:"cost" validate:"required"`
	Description string  `json:"description" db:"description" validate:"required"`
	CategoryID  int     `json:"categoryId" db:"category_id" validate:"required"`
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
