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

func (eor *EventOrganizersRepository) GetTableName() string {
	return "event_organizers"
}

func (eor *EventOrganizersRepository) Find() ([]*models.EventOrganizer, error) {
	var eventsOrganizers []*models.EventOrganizer

	err := eor.goqu.
		From(eor.GetTableName()).ScanStructs(&eventsOrganizers)

	if err != nil {
		log.Fatal(err)
	}

	return eventsOrganizers, nil
}
