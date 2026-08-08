package designsystem

import "github.com/go-chi/chi/v5"

func SetupRoutes(router chi.Router) error {
	handlers := NewHandlers()

	router.Get("/design-system", handlers.DesignSystemPage)

	return nil
}
