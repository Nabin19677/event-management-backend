package repositories

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.io/anilk/crane/models"
)

type EventSessionRepository struct {
	db   *sql.DB
	goqu *goqu.Database
}

func InitEventSessionRepository(db *sql.DB, goqu *goqu.Database) *EventSessionRepository {
	return &EventSessionRepository{db: db, goqu: goqu}
}

func (es *EventSessionRepository) GetTableName() string {
	return "event_sessions"
}

func (es *EventSessionRepository) Insert(newEvent models.NewEventSession) (int, error) {
	query := `INSERT INTO ` + es.GetTableName() + ` (event_id, name, start_time, end_time) VALUES ($1, $2, $3, $4) RETURNING session_id`

	var lastInsertID int
	err := es.db.QueryRow(query, newEvent.EventID, newEvent.Name, newEvent.StartTime, newEvent.EndTime).Scan(&lastInsertID)
	if err != nil {
		return -1, err
	}

	return lastInsertID, nil

}
