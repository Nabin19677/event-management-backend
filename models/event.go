package models

type EventFields struct {
	StartDate   *string `json:"startDate,omitempty" db:"start_date" validate:"required"`
	EndDate     *string `json:"endDate,omitempty" db:"end_date" validate:"required"`
	Location    *string `json:"location,omitempty" db:"location" validate:"required"`
	Description *string `json:"description,omitempty" db:"description" validate:"required"`
}

type Event struct {
	EventID int    `json:"eventId" db:"event_id"`
	Name    string `json:"name" db:"name"`
	EventFields
	AdminUserID int `json:"adminUserId" db:"admin_user_id"`
}

type NewEvent struct {
	Name string `json:"name" db:"name" validate:"required"`
	EventFields
	AdminUserID int `json:"adminUserId" db:"admin_user_id" validate:"required"`
}

type UpdateEvent struct {
	EventFields
}
