package designsystem

import (
	"net/http"

	"github.com/jfkonecn/gophitect/internal/features/designsystem/pages"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) DesignSystemPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.DesignSystemPage().Render(r.Context(), w); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
