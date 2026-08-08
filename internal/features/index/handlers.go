package index

import (
	"net/http"

	"github.com/jfkonecn/gophitect/internal/features/index/pages"
)

type Handlers struct {
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) IndexPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.IndexPage().Render(r.Context(), w); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
