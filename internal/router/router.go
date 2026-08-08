package router

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/jfkonecn/gophitect/internal/config"
	counterFeature "github.com/jfkonecn/gophitect/internal/features/counter"
	indexFeature "github.com/jfkonecn/gophitect/internal/features/index"
	monitorFeature "github.com/jfkonecn/gophitect/internal/features/monitor"
	reverseFeature "github.com/jfkonecn/gophitect/internal/features/reverse"
	sortableFeature "github.com/jfkonecn/gophitect/internal/features/sortable"
	"github.com/jfkonecn/gophitect/internal/web/resources"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"github.com/starfederation/datastar-go/datastar"
)

func SetupRoutes(ctx context.Context, router chi.Router, sessionStore *sessions.CookieStore) (err error) {

	if config.Global.Environment == config.Dev {
		setupReload(router)
	}

	router.Handle("/static/*", resources.Handler())

	if err := errors.Join(
		indexFeature.SetupRoutes(ctx, router, sessionStore),
		counterFeature.SetupRoutes(router, sessionStore),
		monitorFeature.SetupRoutes(router),
		sortableFeature.SetupRoutes(router),
		reverseFeature.SetupRoutes(router),
	); err != nil {
		return fmt.Errorf("error setting up routes: %w", err)
	}

	return nil
}

func setupReload(router chi.Router) {
	reloadChan := make(chan struct{}, 1)

	router.Get("/reload", func(w http.ResponseWriter, r *http.Request) {
		sse := datastar.NewSSE(w, r)
		select {
		case <-reloadChan:
			sse.ExecuteScript("window.location.reload()")
		case <-r.Context().Done():
		}
	})

	router.Get("/hotreload", func(w http.ResponseWriter, r *http.Request) {
		select {
		case reloadChan <- struct{}{}:
		default:
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

}
