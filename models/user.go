package models

type User struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
}

type NewUser struct {
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
}
