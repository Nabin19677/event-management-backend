package repositories

import (
	"database/sql"
	"log"

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

func (es *EventSessionRepository) FindAllByEventId(eventId int) ([]*models.EventSession, error) {
	var eventSessions []*models.EventSession

	err := es.goqu.
		From(es.GetTableName()).Where(goqu.Ex{"event_id": eventId}).ScanStructs(&eventSessions)

	if err != nil {
		log.Println("find failed :", err)
		return eventSessions, err
	}

	return eventSessions, nil
}

func (es *EventSessionRepository) Insert(newEvent models.NewEventSession) (int, error) {
	query, _, _ := goqu.Insert(es.GetTableName()).Rows(newEvent).Returning("session_id").ToSQL()

	var lastInsertID int
	err := es.db.QueryRow(query).Scan(&lastInsertID)
	if err != nil {
		return -1, err
	}

	return lastInsertID, nil

}
