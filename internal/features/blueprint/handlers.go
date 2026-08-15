package blueprint

import (
	"net/http"

	"github.com/jfkonecn/gophitect/internal/features/blueprint/pages"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) Blueprint(w http.ResponseWriter, r *http.Request) {
	if err := pages.Blueprint().Render(r.Context(), w); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
