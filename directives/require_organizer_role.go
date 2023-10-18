package directives

import (
	"context"
	"errors"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.io/anilk/crane/database/postgres/repositories"
	"github.io/anilk/crane/middleware"
	"github.io/anilk/crane/models"
)

func RequireOrganizerRole(eventOrganizersRepository *repositories.EventOrganizersRepository, eventAttendeeRepository *repositories.EventAttendeeRepository) func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []models.Role) (interface{}, error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []models.Role) (interface{}, error) {
		eventId, ok := graphql.GetFieldContext(ctx).Args["eventId"].(int)

		if !ok {
			log.Println("field 'eventId' is required to access event organizer specific routes. please add eventId to field")
			return nil, errors.New("unauthenticated")
		}

		user, _ := middleware.GetCurrentUserFromCTX(ctx)

		role, _ := eventOrganizersRepository.GetEventRole(eventId, user.UserID)

		for _, requiredRole := range roles {
			// User's role matches one of the required roles
			if role == requiredRole.String() {
				return next(ctx)
			}
			if requiredRole.String() == models.RoleAttendee.String() { // Check if "Attendee" role is present in required role

				// Check Attendees
				eventAttendee, _ := eventAttendeeRepository.FindByEventAndUserId(eventId, user.UserID)

				if eventAttendee.EventAttendeeID != 0 {
					return next(ctx)
				}

			}
		}

		return nil, errors.New("unauthorized")
	}
}
