package timeline

import (
	"net/http"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	tu usecase.Timeline
}

func NewRouter(tu usecase.Timeline) http.Handler {
	r := chi.NewRouter()
	h := &handler{tu}

	r.Get("/public", h.Public)

	return r
}
