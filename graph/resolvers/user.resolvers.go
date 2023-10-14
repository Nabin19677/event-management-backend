package resolvers

import (
	"context"
	"fmt"

	"github.io/anilk/crane/models"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (bool, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	users, err := r.UserRepository.Find()
	if err != nil {
		return nil, err
	}
	return users, nil
}
