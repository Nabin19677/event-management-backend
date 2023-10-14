package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.io/anilk/crane/conf"
	"github.io/anilk/crane/database/postgres"
	"github.io/anilk/crane/database/postgres/repositories"
	"github.io/anilk/crane/graph"
	"github.io/anilk/crane/graph/resolvers"
)

const defaultPort = "8080"

func main() {
	conf.InitEnvConfigs()

	db, goqu := postgres.CreateDBConnection()

	userRepository := repositories.InitUserRepository(db, goqu)
	eventRepository := repositories.InitEventRepository(db, goqu)

	resolversMap := &resolvers.Resolver{UserRepository: userRepository, EventRepository: eventRepository}

	port := conf.EnvConfigs.ServerPort
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: resolversMap}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
