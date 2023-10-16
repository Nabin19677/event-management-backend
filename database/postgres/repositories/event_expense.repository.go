package repositories

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.io/anilk/crane/models"
)

type EventExpenseRepository struct {
	db   *sql.DB
	goqu *goqu.Database
}

func InitEventExpenseRepository(db *sql.DB, goqu *goqu.Database) *EventExpenseRepository {
	return &EventExpenseRepository{db: db, goqu: goqu}
}

func (er *EventExpenseRepository) GetTableName() string {
	return "expenses"
}

func (er *EventExpenseRepository) Insert(newEvent models.NewEventExpense) (int, error) {
	query := `INSERT INTO ` + er.GetTableName() + ` (event_id, item_name, cost, description, category_id) VALUES ($1, $2, $3, $4, $5) RETURNING expense_id`

	var lastInsertID int
	err := er.db.QueryRow(query, newEvent.EventID, newEvent.ItemName, newEvent.Cost, newEvent.Description, newEvent.CategoryID).Scan(&lastInsertID)
	if err != nil {
		return -1, err
	}

	return lastInsertID, nil
}

func (er *EventExpenseRepository) GetTotalExpensesByCategory(eventId int) ([]*models.CategoryTotal, error) {

	return nil, nil
}
