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

func (ur *UserRepository) GetTableName() string {
	return "users"
}

// FindByID retrieves a user by their ID.
func (ur *UserRepository) FindByID(userID int) (*models.PublicUser, error) {
	var user models.PublicUser
	_, err := ur.goqu.
		From(ur.GetTableName()).
		Where(goqu.Ex{"user_id": userID}).
		ScanStruct(&user)

	return &user, err
}

// FindByEmail retrieves a user by their Email.
func (ur *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	_, err := ur.goqu.
		From(ur.GetTableName()).
		Where(goqu.Ex{"email": email}).
		ScanStruct(&user)

	return &user, err
}

func (ur *UserRepository) Find() ([]*models.PublicUser, error) {
	var users []*models.PublicUser

	err := ur.goqu.
		From(ur.GetTableName()).ScanStructs(&users)

	if err != nil {
		log.Println("find failed :", err)
		return users, err
	}

	return users, nil
}

func (ur *UserRepository) Insert(newUser models.NewUser) (bool, error) {
	newUser.HashPassword(newUser.Password)
	_, err := ur.goqu.Insert(ur.GetTableName()).Rows(
		newUser,
	).Executor().Exec()

	if err != nil {
		log.Println("insert failed :", err)
		return false, err
	}

	return true, nil

}
