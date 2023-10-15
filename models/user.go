package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	UserID      int     `json:"userId" db:"user_id"`
	Name        string  `json:"name" db:"name"`
	Email       *string `json:"email" db:"email" `
	PhoneNumber *string `json:"phoneNumber" db:"phone_number"`
	Password    string  `json:"password" db:"password"`
}

type PublicUser struct {
	UserID      int     `json:"userId" db:"user_id"`
	Name        string  `json:"name" db:"name"`
	Email       *string `json:"email" db:"email" `
	PhoneNumber *string `json:"phoneNumber" db:"phone_number"`
}

type NewUser struct {
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
	Password    string `json:"password" db:"password"`
}

func (u *NewUser) HashPassword(password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	u.Password = string(passwordHash)

	return err
}
