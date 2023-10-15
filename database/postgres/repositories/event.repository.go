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

func (er *EventRepository) GetTableName() string {
	return "events"
}

func (er *EventRepository) FindByID(eventID int) (*models.Event, error) {
	var event models.Event
	_, err := er.goqu.
		From(er.GetTableName()).
		Where(goqu.Ex{"event_id": eventID}).
		ScanStruct(&event)

	return &event, err
}

func (er *EventRepository) Find() ([]*models.Event, error) {
	var events []*models.Event

	err := er.goqu.
		From(er.GetTableName()).ScanStructs(&events)

	if err != nil {
		log.Fatal(err)
	}

	return events, nil
}

func (er *EventRepository) Insert(newEvent models.NewEvent) (bool, error) {
	_, err := er.goqu.Insert(er.GetTableName()).Rows(
		newEvent,
	).Executor().Exec()

	if err != nil {
		log.Fatal(err)
	}

	return true, nil

}
