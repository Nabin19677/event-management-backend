package repositories

import (
	"database/sql"
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.io/anilk/crane/models"
)

type UserRepository struct {
	db   *sql.DB
	goqu *goqu.Database
}

func InitUserRepository(db *sql.DB, goqu *goqu.Database) *UserRepository {
	return &UserRepository{db: db, goqu: goqu}
}

// FindByID retrieves a user by their ID.
func (ur *UserRepository) FindByID(userID int) (*models.User, error) {
	var user models.User
	_, err := ur.goqu.
		From("users").
		Where(goqu.Ex{"user_id": userID}).
		ScanStruct(&user)

	return &user, err
}

func (ur *UserRepository) Find() ([]*models.User, error) {
	var users []*models.User

	err := ur.goqu.
		From("users").ScanStructs(&users)

	if err != nil {
		log.Fatal(err)
	}

	return users, nil
}

func (ur *UserRepository) Insert(newUser models.NewUser) (bool, error) {
	_, err := ur.goqu.Insert("users").Rows(
		newUser,
	).Executor().Exec()

	if err != nil {
		log.Fatal(err)
	}

	return true, nil
}
