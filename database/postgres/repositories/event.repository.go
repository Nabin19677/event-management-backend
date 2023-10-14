package repositories

import (
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.io/anilk/crane/models"
)

type EventRepository struct {
	db   *sql.DB
	goqu *goqu.Database
}

func InitEventRepository(db *sql.DB, goqu *goqu.Database) *EventRepository {
	return &EventRepository{db: db, goqu: goqu}
}

func (er *EventRepository) Find() ([]*models.Event, error) {
	var events []*models.Event

	err := er.goqu.
		From("events").ScanStructs(&events)

	if err != nil {
		log.Fatal(err)
	}

	return events, nil
}

func (er *EventRepository) Insert(newEvent models.NewEvent) (bool, error) {
	_, err := er.goqu.Insert("events").Rows(
		newEvent,
	).Executor().Exec()

	if err != nil {
		log.Fatal(err)
	}

	return true, nil

}
