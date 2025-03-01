package films

import (
	"net/http"
)

type HandlerFilms struct{}

type HandlerFilmsDeps struct{}

func New(router *http.ServeMux, deps *HandlerFilmsDeps) {
	handler := &HandlerFilms{}

	router.HandleFunc("POST /film", handler.Create())
}

func (h *HandlerFilms) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
