package main

import (
	"log"
	"myapp/graph"
	"myapp/graph/generated"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/LouisAldorio/Testing-early-injection-directive/directives"
	"github.com/LouisAldorio/Testing-early-injection-directive/middleware"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8085"

func main() {

	// service.AddFriend([]string{"louisaldorio","temari","sakura"},"giselia")

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(middleware.Auth())

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.HasRole = directives.HasRole

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	handler := cors.New(cors.Options{
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
	}).Handler(srv)

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", handler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
