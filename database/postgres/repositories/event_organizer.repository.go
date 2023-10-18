package repositories

import (
	"database/sql"
	"errors"
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

func (eor *EventOrganizersRepository) FindByID(eventOrganizerId int) (*models.EventOrganizer, error) {
	var eventOrganizer models.EventOrganizer
	_, err := eor.goqu.
		From(eor.GetTableName()).
		Where(goqu.Ex{"event_organizer_id": eventOrganizerId}).
		ScanStruct(&eventOrganizer)

	return &eventOrganizer, err
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

func (eor *EventOrganizersRepository) Insert(newEventOrganizer models.NewEventOrganizer) (bool, error) {
	_, err := eor.goqu.Insert(eor.GetTableName()).Rows(
		newEventOrganizer,
	).Executor().Exec()

	if err != nil {
		log.Println(err)
		return false, errors.New("could not add event organizer")
	}

	return true, nil
}

func (eor *EventOrganizersRepository) Delete(eventOrganizerId int) (bool, error) {
	_, err := eor.goqu.Delete(eor.GetTableName()).Where(goqu.Ex{"event_organizer_id": eventOrganizerId}).Executor().Exec()

	if err != nil {
		log.Fatal(err)
		return false, err
	}

	return true, nil
}

func (eor *EventOrganizersRepository) GetEventRole(eventId int, userId int) (string, error) {
	var eventOrganizer models.EventOrganizer
	var role models.EventRole
	_, err := eor.goqu.
		From(eor.GetTableName()).
		Where(goqu.Ex{"user_id": userId, "event_id": eventId}).
		ScanStruct(&eventOrganizer)
	_, err = eor.goqu.From("event_roles").Where(goqu.Ex{"role_id": eventOrganizer.RoleID}).ScanStruct(&role)

	return role.RoleName, err
}

func (eor *EventOrganizersRepository) FindByEventId(eventId int) ([]*models.EventOrganizer, error) {
	var eventsOrganizers []*models.EventOrganizer

	err := eor.goqu.
		From(eor.GetTableName()).Where(goqu.Ex{"event_id": eventId}).ScanStructs(&eventsOrganizers)

	if err != nil {
		log.Fatal(err)
	}

	return eventsOrganizers, nil
}
