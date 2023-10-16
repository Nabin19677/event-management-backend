package models

type EventSession struct {
	SessionID int    `json:"sessionId" db:"session_id"`
	EventID   int    `json:"eventId" db:"event_id"`
	Name      string `json:"name" db:"name"`
	StartTime string `json:"startTime" db:"start_time"`
	EndTime   string `json:"endTime" db:"end_time"`
}
