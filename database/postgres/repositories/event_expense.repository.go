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
	query, _, _ := goqu.Insert(er.GetTableName()).Rows(newEvent).Returning("expense_id").ToSQL()

	var lastInsertID int
	err := er.db.QueryRow(query).Scan(&lastInsertID)
	if err != nil {
		return -1, err
	}

	return lastInsertID, nil
}

func (er *EventExpenseRepository) GetTotalExpensesByCategory(eventId int) ([]*models.CategoryTotal, error) {
	// Create a Goqu instance for the expenses table
	expensesTable := er.goqu.From("expenses").Select(
		goqu.I("category_id").As("category_id"),
		goqu.I("cost").As("cost"),
	).Where(
		goqu.C("event_id").Eq(eventId),
	)

	// Join with event_expense_categories to get category names
	query := expensesTable.LeftJoin(
		goqu.T("event_expense_categories"),
		goqu.On(goqu.I("expenses.category_id").Eq(goqu.I("event_expense_categories.category_id"))),
	).Select(
		goqu.I("event_expense_categories.category_name").As("category_name"),
		goqu.SUM(goqu.I("cost")).As("total_cost"),
	).GroupBy("category_name")

	// Execute the query and retrieve the results
	var results []*models.CategoryTotal
	err := query.ScanStructs(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
