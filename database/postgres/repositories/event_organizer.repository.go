package repositories

import (
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.io/anilk/crane/models"
)

type EventOrganizersRepository struct {
	db   *sql.DB
	goqu *goqu.Database
}

func InitEventOrganizersRepository(db *sql.DB, goqu *goqu.Database) *EventOrganizersRepository {
	return &EventOrganizersRepository{db: db, goqu: goqu}
}

func (er *EventOrganizersRepository) Find() ([]*models.EventOrganizer, error) {
	var eventsOrganizers []*models.EventOrganizer

	err := er.goqu.
		From("event_organizers").ScanStructs(&eventsOrganizers)

	if err != nil {
		log.Fatal(err)
	}

	return eventsOrganizers, nil
}
