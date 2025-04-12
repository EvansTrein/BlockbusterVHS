package server

import (
	"net/http"

	httpAdapter "github.com/EvansTrein/BlockbusterVHS/internal/adapters/http"
)

// Grouping of used handlers
type ActiveHandlers struct {
	*httpAdapter.HandlerUser
}

// Grouping of used Middlewares
type ActiveMiddlewares struct {
}

func InitRouters(router *http.ServeMux, handlers *ActiveHandlers, middlewares *ActiveMiddlewares) {
	// Users
	router.Handle("POST /user", handlers.UserCreate())
	router.Handle("GET /user", handlers.User())
	router.Handle("PUT /user", handlers.UserUpd())
	router.Handle("DELETE /user", handlers.UserDel())
}
