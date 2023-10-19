package models

type EventRole struct {
	RoleID   int    `json:"roleId" db:"role_id"`
	RoleName string `json:"roleName" db:"role_name"`
}

type EventOrganizer struct {
	EventOrganizerID int `json:"eventOrganizerId" db:"event_organizer_id"`
	EventID          int `json:"eventId" db:"event_id"`
	UserID           int `json:"userId" db:"user_id"`
	RoleID           int `json:"roleId" db:"role_id"`
}

type NewEventOrganizer struct {
	EventID int `json:"eventId" db:"event_id" validate:"required"`
	UserID  int `json:"userId" db:"user_id" validate:"required"`
	RoleID  int `json:"roleId" db:"role_id" validate:"required"`
}
