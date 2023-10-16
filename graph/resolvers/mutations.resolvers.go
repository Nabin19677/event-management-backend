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
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Println(err)
		return false, errors.New("unauthenticated")
	}
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
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Println(err)
		return false, errors.New("unauthenticated")
	}

	roleId, err := r.EventOrganizersRepository.GetEventRole(eventID, user.UserID)

	if roleId != 1 {
		return false, errors.New("only admin can update event")
	}

	isUpdated, err := r.EventRepository.Update(eventID, input)
	if err != nil {
		return false, err
	}
	return isUpdated, nil
}

// CreateEventOrganizer is the resolver for the createEventOrganizer field.
func (r *mutationResolver) CreateEventOrganizer(ctx context.Context, input models.NewEventOrganizer) (bool, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Println(err)
		return false, errors.New("unauthenticated")
	}

	roleId, err := r.EventOrganizersRepository.GetEventRole(input.EventID, user.UserID)

	if roleId != 1 {
		return false, errors.New("only admin can add event organizers")
	}

	eventOrganizerCreated, err := r.EventOrganizersRepository.Insert(input)
	if err != nil {
		return false, err
	}
	return eventOrganizerCreated, nil
}

// CreateEventAttendee is the resolver for the createEventAttendee field.
func (r *mutationResolver) CreateEventAttendee(ctx context.Context, input models.NewEventAttendee) (bool, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Println(err)
		return false, errors.New("unauthenticated")
	}
	roleId, err := r.EventOrganizersRepository.GetEventRole(input.EventID, user.UserID)

	if roleId != 3 {
		return false, errors.New("only contibutors can add/invite attendees")
	}

	_, err = r.EventAttendeeRepository.Insert(input)
	if err != nil {
		return false, err
	}
	return true, nil
}

// DeleteEventOrganizer is the resolver for the deleteEventOrganizer field.
func (r *mutationResolver) DeleteEventOrganizer(ctx context.Context, eventOrganizerID int) (bool, error) {
	isDeleted, err := r.EventOrganizersRepository.Delete(eventOrganizerID)
	if err != nil {
		return false, err
	}
	return isDeleted, nil
}

// CreateEventSesssion is the resolver for the createEventSesssion field.
func (r *mutationResolver) CreateEventSesssion(ctx context.Context, input models.NewEventSession) (bool, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Println(err)
		return false, errors.New("unauthenticated")
	}
	roleId, err := r.EventOrganizersRepository.GetEventRole(input.EventID, user.UserID)

	if roleId != 1 && roleId != 2 {
		return false, errors.New("only admin/contributors can manage sessions")
	}

	_, err = r.EventSessionRepository.Insert(input)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetEventDetail is the resolver for the getEventDetail field.
func (r *mutationResolver) GetEventDetail(ctx context.Context, eventID int) (*models.EventDetail, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Println(err)
		return nil, errors.New("unauthenticated")
	}

	roleId, err := r.EventOrganizersRepository.GetEventRole(eventID, user.UserID)

	if roleId != 3 {
		return nil, errors.New("are not invited to the event")
	}

	event, err := r.EventRepository.FindByID(eventID)

	sessions, err := r.EventSessionRepository.FindAllByEventId(eventID)

	if err != nil {
		log.Println(err)
		return nil, errors.New("unable to find event details")
	}

	return &models.EventDetail{
		Event:    event,
		Sessions: sessions,
	}, nil
}

// CreateEventExpense is the resolver for the createEventExpense field.
func (r *mutationResolver) CreateEventExpense(ctx context.Context, input models.NewEventExpense) (bool, error) {
	user, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Println(err)
		return false, errors.New("unauthenticated")
	}

	roleId, err := r.EventOrganizersRepository.GetEventRole(input.EventID, user.UserID)

	if roleId != 1 {
		return false, errors.New("cannot add event expenses")
	}

	_, err = r.EventExpenseRepository.Insert(input)

	if err != nil {
		return false, err
	}
	return true, nil
}

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
