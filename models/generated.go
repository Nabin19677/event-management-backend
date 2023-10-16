// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

type AuthResponse struct {
	AuthToken *AuthToken `json:"authToken"`
}

type AuthToken struct {
	AccessToken string `json:"accessToken"`
	ExpireAt    string `json:"expireAt"`
}

type EventDetail struct {
	Event    *Event          `json:"event,omitempty"`
	Sessions []*EventSession `json:"sessions,omitempty"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewEventAttendee struct {
	EventID int `json:"eventId"`
	UserID  int `json:"userId"`
}

type NewEventExpense struct {
	EventID     int     `json:"eventId"`
	ItemName    string  `json:"itemName"`
	Cost        float64 `json:"cost"`
	Description *string `json:"description,omitempty"`
	CategoryID  int     `json:"categoryId"`
}

type NewEventSession struct {
	EventID   int    `json:"eventId"`
	Name      string `json:"name"`
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type UpdateEvent struct {
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
	Location    string `json:"location"`
	Description string `json:"description"`
}
