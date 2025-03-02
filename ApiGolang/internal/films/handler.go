package films

import (
	"log/slog"
	"net/http"
)

type HandlerFilms struct {
	log *slog.Logger
}

type HandlerFilmsDeps struct {
	*slog.Logger
}

func NewHandler(router *http.ServeMux, deps *HandlerFilmsDeps) {
	handler := &HandlerFilms{
		log: deps.Logger,
	}

	router.HandleFunc("POST /film", handler.Create())
}

func (h *HandlerFilms) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		op := "Handler Create: call"
		log := h.log.With(
			slog.String("operation", op),
			slog.String("apiPath", r.URL.Path),
			slog.String("HTTP Method", r.Method),
		)
		log.Debug("request received")
	}
}
