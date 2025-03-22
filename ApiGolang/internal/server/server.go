package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/EvansTrein/BlockbusterVHS/config"
	"github.com/EvansTrein/BlockbusterVHS/internal/films"
	"github.com/EvansTrein/BlockbusterVHS/internal/storages/postgres"
	"github.com/EvansTrein/BlockbusterVHS/internal/users"
	"github.com/EvansTrein/BlockbusterVHS/pkg/middleware"
)

const (
	gracefulShutdownTimer = time.Second * 10
)

type HttpServer struct {
	conf   *config.HTTPServer
	log    *slog.Logger
	server *http.Server
	router *http.ServeMux
}

type HttpServerDeps struct {
	*config.HTTPServer
	*slog.Logger
	*postgres.PostgresDB
}

func New(deps *HttpServerDeps) *HttpServer {
	router := http.NewServeMux()

	// Repositories
	repoUsers := users.NewUsersRepoPostgres(&users.UsersRepoPostgresDeps{
		Logger:     deps.Logger,
		PostgresDB: deps.PostgresDB,
	})

	// Services
	serviceUsers := users.NewUsersService(&users.UsersServiceDeps{
		Logger:     deps.Logger,
		IUsersRepo: repoUsers,
	})

	// Handlers
	films.NewHandler(router, &films.HandlerFilmsDeps{
		Logger: deps.Logger,
	})

	users.NewHandler(router, &users.HandlerUsersDeps{
		Logger:        deps.Logger,
		IUsersService: serviceUsers,
	})

	return &HttpServer{
		conf:   deps.HTTPServer,
		log:    deps.Logger,
		router: router,
	}
}

func (s *HttpServer) Start() error {
	log := s.log.With(slog.String("Address", s.conf.Address+":"+s.conf.Port))
	log.Debug("HTTP server: started creating")

	s.server = &http.Server{
		Addr: s.conf.Address + ":" + s.conf.Port,
		Handler: middleware.ChainMiddleware(
			middleware.Timeout(s.conf.WriteTimeout),
			middleware.CORS,
		)(s.router),
		ReadHeaderTimeout: s.conf.ReadHeaderTimeout,
		ReadTimeout:       s.conf.ReadTimeout,
		WriteTimeout:      s.conf.WriteTimeout,
		IdleTimeout:       s.conf.IdleTimeout,
	}

	log.Info("HTTP server: successfully started")
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *HttpServer) Stop() error {
	s.log.Debug("HTTP server: stop started")

	ctx, cancel := context.WithTimeout(context.Background(), gracefulShutdownTimer)
	defer cancel()

	if err := s.server.Shutdown(ctx); err != nil {
		s.log.Error("Server shutdown failed", "error", err)
		return err
	}

	s.server = nil
	s.log.Info("HTTP server: stop successful")

	return nil
}
