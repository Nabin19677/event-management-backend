package directives

import (
	"context"
	"errors"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.io/anilk/crane/middleware"
)

func Authenticate(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	_, err = middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		log.Println(err)
		return false, errors.New("unauthenticated")
	}
	return next(ctx)
}
