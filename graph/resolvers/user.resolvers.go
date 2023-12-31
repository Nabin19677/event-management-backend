package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.38

import (
	"context"
	"errors"

	"github.io/anilk/crane/lib/validation"
	"github.io/anilk/crane/models"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input models.NewUser) (bool, error) {
	if err := validation.ValidateStruct(input); err != nil {
		return false, err
	}
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
	if err := validation.ValidateStruct(input); err != nil {
		return nil, err
	}
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
