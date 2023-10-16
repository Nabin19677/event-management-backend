package main

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.io/anilk/crane/conf"
	"github.io/anilk/crane/database/postgres"
	"github.io/anilk/crane/database/postgres/repositories"
	"github.io/anilk/crane/graph"
	"github.io/anilk/crane/graph/resolvers"
	appMiddleware "github.io/anilk/crane/middleware"
	"github.io/anilk/crane/models"
)

const defaultPort = "8080"

func main() {
	conf.InitEnvConfigs()

	db, goqu := postgres.CreateDBConnection()

	userRepository := repositories.InitUserRepository(db, goqu)
	eventRepository := repositories.InitEventRepository(db, goqu)
	eventOrganizersRepository := repositories.InitEventOrganizersRepository(db, goqu)
	eventRoleRepository := repositories.InitEventRoleRepository(db, goqu)
	eventAttendeeRepository := repositories.InitEventAttendeeRepository(db, goqu)
	eventSessionRepository := repositories.InitEventSessionRepository(db, goqu)
	eventExpenseRepository := repositories.InitEventExpenseRepository(db, goqu)
	eventExpenseCategoryRepository := repositories.InitEventExpenseCategoryRepository(db, goqu)

	resolversMap := &resolvers.Resolver{
		UserRepository:                 userRepository,
		EventRepository:                eventRepository,
		EventOrganizersRepository:      eventOrganizersRepository,
		EventRoleRepository:            eventRoleRepository,
		EventAttendeeRepository:        eventAttendeeRepository,
		EventSessionRepository:         eventSessionRepository,
		EventExpenseRepository:         eventExpenseRepository,
		EventExpenseCategoryRepository: eventExpenseCategoryRepository,
	}

	port := conf.EnvConfigs.ServerPort
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Use "*" to allow all origins
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler

	router.Use(corsHandler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(appMiddleware.AuthMiddleware(userRepository))

	c := graph.Config{Resolvers: resolversMap}

	c.Directives.RequireOrganizerRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []models.Role) (interface{}, error) {

		eventId, ok := graphql.GetFieldContext(ctx).Args["eventId"].(int)

		if !ok {
			log.Println("field 'eventId' is required to access event organizer specific routes. please add eventId to field")
			return nil, errors.New("unauthenticated")
		}

		user, _ := appMiddleware.GetCurrentUserFromCTX(ctx)

		role, _ := eventOrganizersRepository.GetEventRole(eventId, user.UserID)

		for _, requiredRole := range roles {
			if role == requiredRole.String() {
				// User's role matches one of the required roles
				return next(ctx)
			}
		}

		return nil, errors.New("unauthorized")
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
