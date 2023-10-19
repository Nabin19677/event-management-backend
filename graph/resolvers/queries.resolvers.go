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

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.PublicUser, error) {
	users, err := r.UserRepository.Find()
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Events is the resolver for the events field.
func (r *queryResolver) Events(ctx context.Context) ([]*models.Event, error) {
	events, err := r.EventRepository.Find()
	if err != nil {
		return nil, err
	}
	return events, nil
}

// EventsOrganizers is the resolver for the events_organizers field.
func (r *queryResolver) EventsOrganizers(ctx context.Context) ([]*models.EventOrganizer, error) {
	eventsOrganizers, err := r.EventOrganizersRepository.Find()
	if err != nil {
		return nil, err
	}
	return eventsOrganizers, nil
}

// EventsRoles is the resolver for the events_roles field.
func (r *queryResolver) EventsRoles(ctx context.Context) ([]*models.EventRole, error) {
	eventsRoles, err := r.EventRoleRepository.Find()
	if err != nil {
		return nil, err
	}
	return eventsRoles, nil
}

// EventsExpenseCategories is the resolver for the events_expense_categories field.
func (r *queryResolver) EventsExpenseCategories(ctx context.Context) ([]*models.EventExpenseCategory, error) {
	categories, err := r.EventExpenseCategoryRepository.Find()
	if err != nil {
		return nil, err
	}
	return categories, nil
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

// GetEventExpensesByCategory is the resolver for the getEventExpensesByCategory field.
func (r *queryResolver) GetEventExpensesByCategory(ctx context.Context, eventID int) ([]*models.CategoryTotal, error) {
	totalExpenses, _ := r.EventExpenseRepository.GetTotalExpensesByCategory(eventID)

	return totalExpenses, nil
}

// GetEventOrganizers is the resolver for the getEventOrganizers field.
func (r *queryResolver) GetEventOrganizers(ctx context.Context, eventID int) ([]*models.EventOrganizer, error) {
	eventsOrganizers, err := r.EventOrganizersRepository.FindByEventId(eventID)
	if err != nil {
		return nil, err
	}
	return eventsOrganizers, nil
}

// GetEventSessions is the resolver for the getEventSessions field.
func (r *queryResolver) GetEventSessions(ctx context.Context, eventID int) ([]*models.EventSession, error) {
	eventsOrganizers, err := r.EventSessionRepository.FindAllByEventId(eventID)
	if err != nil {
		return nil, err
	}
	return eventsOrganizers, nil
}

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
