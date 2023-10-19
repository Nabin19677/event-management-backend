package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"errors"

	"github.io/anilk/crane/graph"
	"github.io/anilk/crane/models"
)

// EventID is the resolver for the eventId field.
func (r *eventOrganizerResolver) EventID(ctx context.Context, obj *models.EventOrganizer) (*models.Event, error) {
	event, err := r.EventRepository.FindByID(obj.EventID)

	if err != nil {
		return nil, err
	}

	return event, nil
}

// UserID is the resolver for the userId field.
func (r *eventOrganizerResolver) UserID(ctx context.Context, obj *models.EventOrganizer) (*models.PublicUser, error) {
	user, err := r.UserRepository.FindByIDPublic(obj.UserID)

	if err != nil {
		return nil, err
	}

	return user, nil
}

// RoleID is the resolver for the roleId field.
func (r *eventOrganizerResolver) RoleID(ctx context.Context, obj *models.EventOrganizer) (*models.EventRole, error) {
	role, err := r.EventRoleRepository.FindByID(obj.RoleID)

	if err != nil {
		return nil, err
	}

	return role, nil
}

// CreateEventOrganizer is the resolver for the createEventOrganizer field.
func (r *mutationResolver) CreateEventOrganizer(ctx context.Context, eventID int, input models.NewEventOrganizer) (bool, error) {
	eventOrganizerCreated, err := r.EventOrganizersRepository.Insert(input)
	if err != nil {
		return false, err
	}
	// Check If Role Attendee
	if input.RoleID == 3 {
		_, err := r.EventAttendeeRepository.Insert(models.NewEventAttendee{
			EventID: eventID,
			UserID:  input.UserID,
		})
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return eventOrganizerCreated, nil
}

// DeleteEventOrganizer is the resolver for the deleteEventOrganizer field.
func (r *mutationResolver) DeleteEventOrganizer(ctx context.Context, eventID int, eventOrganizerID int) (bool, error) {
	event, err := r.EventRepository.FindByID(eventID)
	eventOrganizer, err := r.EventOrganizersRepository.FindByID(eventOrganizerID)

	if event.AdminUserID == eventOrganizer.UserID {
		return false, errors.New("cannot delete creator of the event.")
	}

	isDeleted, err := r.EventOrganizersRepository.Delete(eventOrganizerID)
	if err != nil {
		return false, err
	}
	return isDeleted, nil
}

// GetEventOrganizers is the resolver for the getEventOrganizers field.
func (r *queryResolver) GetEventOrganizers(ctx context.Context, eventID int) ([]*models.EventOrganizer, error) {
	eventsOrganizers, err := r.EventOrganizersRepository.FindByEventId(eventID)
	if err != nil {
		return nil, err
	}
	return eventsOrganizers, nil
}

// EventOrganizer returns graph.EventOrganizerResolver implementation.
func (r *Resolver) EventOrganizer() graph.EventOrganizerResolver { return &eventOrganizerResolver{r} }

type eventOrganizerResolver struct{ *Resolver }
