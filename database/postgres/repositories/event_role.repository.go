package repositories

import (
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.io/anilk/crane/models"
)

type EventRoleRepository struct {
	db   *sql.DB
	goqu *goqu.Database
}

func InitEventRoleRepository(db *sql.DB, goqu *goqu.Database) *EventRoleRepository {
	return &EventRoleRepository{db: db, goqu: goqu}
}

func (er *EventRoleRepository) FindByID(roleID int) (*models.EventRole, error) {
	var role models.EventRole
	_, err := er.goqu.
		From("event_roles").
		Where(goqu.Ex{"role_id": roleID}).
		ScanStruct(&role)

	return &role, err
}

func (er *EventRoleRepository) Find() ([]*models.EventRole, error) {
	var roles []*models.EventRole

	err := er.goqu.
		From("event_roles").ScanStructs(&roles)

	if err != nil {
		log.Fatal(err)
	}

	return roles, nil
}
