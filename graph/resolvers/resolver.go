//go:generate go run github.com/99designs/gqlgen@v0.17.38 generate
package resolvers

import "github.io/anilk/crane/database/postgres/repositories"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserRepository            *repositories.UserRepository
	EventRepository           *repositories.EventRepository
	EventOrganizersRepository *repositories.EventOrganizersRepository
	EventRoleRepository       *repositories.EventRoleRepository
	EventAttendeeRepository   *repositories.EventAttendeeRepository
}
