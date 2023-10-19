package models

type EventSessionFields struct {
	EventID   int    `json:"eventId" db:"event_id" validate:"required"`
	Name      string `json:"name" db:"name" validate:"required"`
	StartTime string `json:"startTime" db:"start_time" validate:"required"`
	EndTime   string `json:"endTime" db:"end_time" validate:"required"`
}
type EventSession struct {
	SessionID int `json:"sessionId" db:"session_id"`
	EventSessionFields
}

type NewEventSession struct {
	EventSessionFields
}
