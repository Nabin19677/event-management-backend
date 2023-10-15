package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"

	"github.io/anilk/crane/graph"
	"github.io/anilk/crane/models"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (bool, error) {
	user, err := r.UserRepository.Insert(input)
	if err != nil {
		return false, err
	}
	return user, nil
}

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input models.NewEvent) (bool, error) {
	event, err := r.EventRepository.Insert(input)
	if err != nil {
		return false, err
	}
	return event, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
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

// Mutation returns graph.MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }