package repositories

import (
	"database/sql"

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
