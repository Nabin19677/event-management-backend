package repositories

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	"github.io/anilk/crane/models"
)

type UserRepository struct {
	db           *sqlx.DB
	queryBuilder *goqu.Database
}

func InitUserRepository(db *sqlx.DB, queryBuilder *goqu.Database) *UserRepository {
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
	query, _, _ := ur.queryBuilder.
		From("users").Select("id", "name", "email", "phone_number").ToSQL()

	rows, err := ur.db.Queryx(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*models.User

	for rows.Next() {
		var user models.User

		rows.Scan(&user.ID, &user.Name, &user.Email, &user.PhoneNumber)

		users = append(users, &user)
	}

	return users, nil
}
