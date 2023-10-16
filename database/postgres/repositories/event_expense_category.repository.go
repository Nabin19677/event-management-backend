package repositories

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.io/anilk/crane/models"
)

type EventExpenseCategoryRepository struct {
	db   *sql.DB
	goqu *goqu.Database
}

func InitEventExpenseCategoryRepository(db *sql.DB, goqu *goqu.Database) *EventExpenseCategoryRepository {
	return &EventExpenseCategoryRepository{db: db, goqu: goqu}
}

func (er *EventExpenseCategoryRepository) GetTableName() string {
	return "event_expense_categories"
}

func (er *EventExpenseCategoryRepository) FindByID(categoryID int) (*models.EventExpenseCategory, error) {
	var category models.EventExpenseCategory
	_, err := er.goqu.
		From(er.GetTableName()).
		Where(goqu.Ex{"category_id": categoryID}).
		ScanStruct(&category)

	return &category, err
}
