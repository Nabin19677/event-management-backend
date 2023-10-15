package models

type EventAttendee struct {
	EventAttendeeID int `json:"eventAttendeeId" db:"event_attendee_id"`
	EventID         int `json:"eventId" db:"event_id"`
	UserID          int `json:"userId" db:"user_id"`
}
