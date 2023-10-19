package models

type EventFields struct {
	Name        string  `json:"name" db:"name"`
	StartDate   *string `json:"startDate,omitempty" db:"start_date"`
	EndDate     *string `json:"endDate,omitempty" db:"end_date"`
	Location    *string `json:"location,omitempty" db:"location"`
	Description *string `json:"description,omitempty" db:"description"`
	AdminUserID int     `json:"adminUserId" db:"admin_user_id"`
}

type Event struct {
	EventID int `json:"eventId" db:"event_id"`
	EventFields
}

type NewEvent struct {
	EventFields
}
