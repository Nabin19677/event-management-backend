// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
)

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
	Role     *string         `json:"role,omitempty"`
}

type Role string

const (
	RoleAdmin       Role = "Admin"
	RoleContributor Role = "Contributor"
	RoleAttendee    Role = "Attendee"
)

var AllRole = []Role{
	RoleAdmin,
	RoleContributor,
	RoleAttendee,
}

func (e Role) IsValid() bool {
	switch e {
	case RoleAdmin, RoleContributor, RoleAttendee:
		return true
	}
	return false
}

func (e Role) String() string {
	return string(e)
}

func (e *Role) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Role(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Role", str)
	}
	return nil
}

func (e Role) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
