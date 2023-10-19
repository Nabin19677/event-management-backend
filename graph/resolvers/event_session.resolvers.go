package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"

	"github.io/anilk/crane/graph"
	"github.io/anilk/crane/models"
)

// EventID is the resolver for the eventId field.
func (r *eventSessionResolver) EventID(ctx context.Context, obj *models.EventSession) (*models.Event, error) {
	event, err := r.EventRepository.FindByID(obj.EventID)

	if err != nil {
		return nil, err
	}

	return event, nil
}

// CreateEventSesssion is the resolver for the createEventSesssion field.
func (r *mutationResolver) CreateEventSesssion(ctx context.Context, eventID int, input models.NewEventSession) (bool, error) {
	_, err := r.EventSessionRepository.Insert(input)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetEventSessions is the resolver for the getEventSessions field.
func (r *queryResolver) GetEventSessions(ctx context.Context, eventID int) ([]*models.EventSession, error) {
	eventsOrganizers, err := r.EventSessionRepository.FindAllByEventId(eventID)
	if err != nil {
		return nil, err
	}
	return eventsOrganizers, nil
}

// EventSession returns graph.EventSessionResolver implementation.
func (r *Resolver) EventSession() graph.EventSessionResolver { return &eventSessionResolver{r} }

type eventSessionResolver struct{ *Resolver }