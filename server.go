package main

import (
	"log"
	"net/http"

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

	resolversMap := &resolvers.Resolver{
		UserRepository:            userRepository,
		EventRepository:           eventRepository,
		EventOrganizersRepository: eventOrganizersRepository,
		EventRoleRepository:       eventRoleRepository,
		EventAttendeeRepository:   eventAttendeeRepository,
		EventSessionRepository:    eventSessionRepository,
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

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolversMap}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
