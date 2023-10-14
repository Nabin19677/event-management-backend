package models

type User struct {
	UserID      int     `json:"userId" db:"user_id"`
	Name        string  `json:"name" db:"name"`
	Email       *string `json:"email" db:"email" `
	PhoneNumber *string `json:"phoneNumber" db:"phone_number"`
}

type NewUser struct {
	Name        string `json:"name" db:"name"`
	Email       string `json:"email" db:"email"`
	PhoneNumber string `json:"phoneNumber" db:"phone_number"`
}
