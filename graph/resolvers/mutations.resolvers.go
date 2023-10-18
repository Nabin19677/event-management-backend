package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"errors"

	"github.io/anilk/crane/graph"
	"github.io/anilk/crane/middleware"
	"github.io/anilk/crane/models"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (bool, error) {
	existingUser, err := r.UserRepository.FindByEmail(input.Email)
	if existingUser != nil {
		return false, errors.New("email already in use")
	}
	_, err = r.UserRepository.Insert(input)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := r.UserRepository.FindByEmail(input.Email)
	if err != nil {
		return nil, errors.New("email or password is wrong")
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, errors.New("email or password is wrong")
	}

	token, err := user.GenToken()

	if err != nil {
		return nil, errors.New("something went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
	}, nil
}

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input models.NewEvent) (bool, error) {
	user, _ := middleware.GetCurrentUserFromCTX(ctx)
	input.AdminUserID = user.UserID
	eventId, err := r.EventRepository.Insert(input)
	if err != nil {
		return false, err
	}
	_, err = r.EventOrganizersRepository.Insert(models.NewEventOrganizer{
		EventID: eventId,
		UserID:  user.UserID,
		RoleID:  1,
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// UpdateEvent is the resolver for the updateEvent field.
func (r *mutationResolver) UpdateEvent(ctx context.Context, eventID int, input *models.UpdateEvent) (bool, error) {
	isUpdated, err := r.EventRepository.Update(eventID, input)
	if err != nil {
		return false, err
	}
	return isUpdated, nil
}

// CreateEventOrganizer is the resolver for the createEventOrganizer field.
func (r *mutationResolver) CreateEventOrganizer(ctx context.Context, eventID int, input models.NewEventOrganizer) (bool, error) {
	eventOrganizerCreated, err := r.EventOrganizersRepository.Insert(input)
	if err != nil {
		return false, err
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

// CreateEventSesssion is the resolver for the createEventSesssion field.
func (r *mutationResolver) CreateEventSesssion(ctx context.Context, eventID int, input models.NewEventSession) (bool, error) {
	_, err := r.EventSessionRepository.Insert(input)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateEventAttendee is the resolver for the createEventAttendee field.
func (r *mutationResolver) CreateEventAttendee(ctx context.Context, eventID int, input models.NewEventAttendee) (bool, error) {
	_, err := r.EventAttendeeRepository.Insert(input)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreateEventExpense is the resolver for the createEventExpense field.
func (r *mutationResolver) CreateEventExpense(ctx context.Context, eventID int, input models.NewEventExpense) (bool, error) {
	_, err := r.EventExpenseRepository.Insert(input)

	if err != nil {
		return false, err
	}
	return true, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
