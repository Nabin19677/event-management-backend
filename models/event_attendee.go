package models

type EventAttendeeFields struct {
	EventID int `json:"eventId" db:"event_id"`
	UserID  int `json:"userId" db:"user_id"`
}
type EventAttendee struct {
	EventAttendeeID int `json:"eventAttendeeId" db:"event_attendee_id"`
	EventAttendeeFields
}

type NewEventAttendee struct {
	EventAttendeeFields
}
