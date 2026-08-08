package index

import (
	"net/http"
	"strconv"

	"github.com/jfkonecn/gophitect/internal/features/index/pages"

	"github.com/go-chi/chi/v5"
)

type Handlers struct {
}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) IndexPage(w http.ResponseWriter, r *http.Request) {
	if err := pages.IndexPage("Gophitect").Render(r.Context(), w); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (h *Handlers) parseIndex(w http.ResponseWriter, r *http.Request) (int, error) {
	idx := chi.URLParam(r, "idx")
	i, err := strconv.Atoi(idx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 0, err
	}
	return i, nil
}
