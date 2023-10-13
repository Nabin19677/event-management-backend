package repositories

import (
	"github.com/doug-martin/goqu/v9"
	"github.io/anilk/crane/models"
)

type UserRepository struct {
	db *goqu.Database
}

func InitUserRepository(db *goqu.Database) *UserRepository {
	return &UserRepository{db: db}
}

// FindByID retrieves a user by their ID.
func (ur *UserRepository) FindByID(userID int) (*models.User, error) {
	var user models.User
	_, err := ur.db.
		From("users").
		Where(goqu.Ex{"id": userID}).
		ScanStruct(&user)

	return &user, err
}

func (ur *UserRepository) Find() ([]*models.User, error) {
	var users []*models.User

	err := ur.db.
		From("users").
		ScanStructs(&users)

	if err != nil {
		return nil, err
	}

	return users, nil
}
