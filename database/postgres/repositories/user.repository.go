package repositories

import (
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.io/anilk/crane/models"
)

type UserRepository struct {
	db           *sql.DB
	queryBuilder *goqu.Database
}

func InitUserRepository(db *sql.DB, queryBuilder *goqu.Database) *UserRepository {
	return &UserRepository{db: db, queryBuilder: queryBuilder}
}

// FindByID retrieves a user by their ID.
func (ur *UserRepository) FindByID(userID int) (*models.User, error) {
	var user models.User
	_, err := ur.queryBuilder.
		From("users").
		Where(goqu.Ex{"id": userID}).
		ScanStruct(&user)

	return &user, err
}

func (ur *UserRepository) Find() ([]*models.User, error) {
	var users []*models.User

	err := ur.queryBuilder.
		From("users").ScanStructs(&users)

	if err != nil {
		log.Fatal(err)
	}

	return users, nil
}
