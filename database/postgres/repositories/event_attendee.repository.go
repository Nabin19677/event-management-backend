package repositories

import (
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.io/anilk/crane/models"
)

type EventAttendeeRepository struct {
	db   *sql.DB
	goqu *goqu.Database
}

func InitEventAttendeeRepository(db *sql.DB, goqu *goqu.Database) *EventAttendeeRepository {
	return &EventAttendeeRepository{db: db, goqu: goqu}
}

func (ea *EventAttendeeRepository) GetTableName() string {
	return "event_attendees"
}

func (ea *EventAttendeeRepository) Insert(newEvent models.NewEventAttendee) (int, error) {
	query := `INSERT INTO ` + ea.GetTableName() + ` (event_id, user_id) VALUES ($1, $2) RETURNING event_id`

	var lastInsertID int
	err := ea.db.QueryRow(query, newEvent.EventID, newEvent.UserID).Scan(&lastInsertID)
	if err != nil {
		return -1, err
	}

	return lastInsertID, nil

}

func (ea *EventAttendeeRepository) FindByEventAndUserId(eventId int, userId int) (*models.EventAttendee, error) {
	var eventAttendee models.EventAttendee
	_, err := ea.goqu.
		From(ea.GetTableName()).
		Where(goqu.Ex{"event_id": eventId, "user_id": userId}).
		ScanStruct(&eventAttendee)

	if err != nil {
		return nil, err
	}

	return &eventAttendee, err
}

func (es *EventAttendeeRepository) FindAllByEventId(eventId int) ([]*models.EventAttendee, error) {
	var eventAttendees []*models.EventAttendee

	err := es.goqu.
		From(es.GetTableName()).Where(goqu.Ex{"event_id": eventId}).ScanStructs(&eventAttendees)

	if err != nil {
		log.Println("find failed :", err)
		return eventAttendees, err
	}

	return eventAttendees, nil
}
