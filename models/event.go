package models

type EventFields struct {
	StartDate   *string `json:"startDate,omitempty" db:"start_date"`
	EndDate     *string `json:"endDate,omitempty" db:"end_date"`
	Location    *string `json:"location,omitempty" db:"location"`
	Description *string `json:"description,omitempty" db:"description"`
}

type Event struct {
	EventID int    `json:"eventId" db:"event_id"`
	Name    string `json:"name" db:"name"`
	EventFields
	AdminUserID int `json:"adminUserId" db:"admin_user_id"`
}

type NewEvent struct {
	Name string `json:"name" db:"name"`
	EventFields
	AdminUserID int `json:"adminUserId" db:"admin_user_id"`
}

type UpdateEvent struct {
	EventFields
}
