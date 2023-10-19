package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"errors"
	"log"

	"github.io/anilk/crane/graph"
	"github.io/anilk/crane/middleware"
	"github.io/anilk/crane/models"
)

// AdminUserID is the resolver for the adminUserId field.
func (r *eventResolver) AdminUserID(ctx context.Context, obj *models.Event) (*models.PublicUser, error) {
	user, err := r.UserRepository.FindByIDPublic(obj.AdminUserID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// OrganizedEvents is the resolver for the organized_events field.
func (r *queryResolver) OrganizedEvents(ctx context.Context) ([]*models.Event, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	events, err := r.EventRepository.FindByOrganizerId(user.UserID)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// GetEventDetail is the resolver for the getEventDetail field.
func (r *queryResolver) GetEventDetail(ctx context.Context, eventID int) (*models.EventDetail, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)

	event, err := r.EventRepository.FindByID(eventID)

	role, err := r.EventOrganizersRepository.GetEventRole(eventID, user.UserID)

	sessions, err := r.EventSessionRepository.FindAllByEventId(eventID)

	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to find event details")
	}

	return &models.EventDetail{
		Event:    event,
		Sessions: sessions,
		Role:     &role,
	}, nil
}

// Event returns graph.EventResolver implementation.
func (r *Resolver) Event() graph.EventResolver { return &eventResolver{r} }

type eventResolver struct{ *Resolver }
