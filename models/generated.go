// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
}

type AuthToken struct {
	AccessToken string `json:"accessToken"`
	ExpireAt    string `json:"expireAt"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewEventAttendee struct {
	EventID int `json:"eventId"`
	UserID  int `json:"userId"`
}

type UpdateEvent struct {
	StartDate   *string `json:"startDate,omitempty"`
	EndDate     *string `json:"endDate,omitempty"`
	Location    *string `json:"location,omitempty"`
	Description *string `json:"description,omitempty"`
}
