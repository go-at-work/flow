package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/arisromil/flow/config"
	"github.com/arisromil/flow/domain"
	"github.com/arisromil/flow/graph"
	"github.com/arisromil/flow/jwt"
	"github.com/arisromil/flow/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/tools/playground"
)

func main() {
	ctx := context.Background()

	conf := config.New()

	db := postgres.NewDB(ctx, conf)

	if err := db.Migrate(); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.RequestID)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)
	router.Use(middleware.Timeout(time.Second * 60))

	userRepo := postgres.NewUserRepository(db)
	authTokenService := jwt.NewTokenService(conf)
	authService := domain.NewAuthService(userRepo)

	router.Use(authMiddleware(authTokenService))
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{
				Resolvers: &graph.Resolver{
					AuthService: authService,
				},
			},
		),
	))

	log.Fatal(http.ListenAndServe(":8080", router))

}
