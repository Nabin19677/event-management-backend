package models

type EventOrganizer struct {
	EventOrganizerID int `json:"eventOrganizerId" db:"event_organizer_id"`
	EventID          int `json:"eventId" db:"event_id"`
	UserID           int `json:"userId" db:"user_id"`
	RoleID           int `json:"roleId" db:"role_id"`
}

type NewEventOrganizer struct {
	EventID int `json:"eventId" db:"event_id"`
	UserID  int `json:"userId" db:"user_id"`
	RoleID  int `json:"roleId" db:"role_id"`
}
