package repositories

import (
	"database/sql"
	"errors"
	"log"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/lib/pq"
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

func (er *EventRepository) Insert(newEvent models.NewEvent) (int, error) {
	query := `INSERT INTO ` + er.GetTableName() + ` (name, start_date, end_date, location, description, admin_user_id) VALUES ($1, $2, $3, $4, $5, $6) RETURNING event_id`

	var lastInsertID int
	err := er.db.QueryRow(query, newEvent.Name, newEvent.StartDate, newEvent.EndDate, newEvent.Location, newEvent.Description, newEvent.AdminUserID).Scan(&lastInsertID)
	if err != nil {
		return -1, err
	}

	return lastInsertID, nil

}

func (er *EventRepository) Update(eventID int, updatedEvent *models.UpdateEvent) (bool, error) {
	ds := goqu.Update(er.GetTableName()).Set(
		goqu.Record{
			"start_date":  updatedEvent.StartDate,
			"end_date":    updatedEvent.EndDate,
			"location":    updatedEvent.Location,
			"description": updatedEvent.Description,
		},
	).Where(goqu.Ex{"event_id": eventID})
	updateSQL, _, _ := ds.ToSQL()

	_, err := er.db.Query(updateSQL)

	if err != nil {
		return false, errors.New("update event failed")
	}

	return true, err
}

func (er *EventRepository) FindByOrganizerId(userId int) ([]*models.Event, error) {
	var events []*models.Event

	err := er.goqu.
		From(er.GetTableName()).ScanStructs(&events)

	if err != nil {
		log.Fatal(err)
	}

	return events, nil
}
