package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/elisalimli/meetmeup/common"
	"github.com/elisalimli/meetmeup/graphql/generated"
	"github.com/elisalimli/meetmeup/graphql/resolvers"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

func main() {
	// ...
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// ...

	db, err := common.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	// DB.AddQueryHook(postgres.DBLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	config := generated.Config{Resolvers: &resolvers.Resolver{}}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	customCtx := &common.CustomContext{
		Database: db,
	}

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", common.CreateContext(customCtx, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
