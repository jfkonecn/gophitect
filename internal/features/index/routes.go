package index

import (
	"context"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
)

func SetupRoutes(ctx context.Context, router chi.Router, store sessions.Store) error {
	handlers := NewHandlers()

	router.Get("/", handlers.IndexPage)

	router.Route("/api", func(apiRouter chi.Router) {
	})

	return nil
}
