package blueprint

import "github.com/go-chi/chi/v5"

func SetupRoutes(router chi.Router) error {
	handlers := NewHandlers()

	router.Get("/blueprint", handlers.Blueprint)

	return nil
}
