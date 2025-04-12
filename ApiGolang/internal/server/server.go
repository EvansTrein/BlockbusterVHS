package server

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/EvansTrein/BlockbusterVHS/config"
	httpAdapter "github.com/EvansTrein/BlockbusterVHS/internal/adapters/http"
	"github.com/EvansTrein/BlockbusterVHS/internal/adapters/repository/userRepo"
	useCase "github.com/EvansTrein/BlockbusterVHS/internal/usecase"
	"github.com/EvansTrein/BlockbusterVHS/pkg/db/sqlite"
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
	*sqlite.SqliteDB
}

func New(deps *HttpServerDeps) *HttpServer {
	router := http.NewServeMux()

	// Initialize repository
	repoUsers := userRepo.NewUsersRepo(&userRepo.UsersRepoDeps{
		Logger:   deps.Logger,
		SqliteDB: deps.SqliteDB,
	})

	// Initialize use case
	userUC := useCase.NewUserUseCase(&useCase.UserUseCaseDeps{
		Logger:          deps.Logger,
		IUserRepository: repoUsers,
	})

	// Initialize handler
	baseHandler := httpAdapter.NewBaseHandler(&httpAdapter.BaseHandlerDeps{
		Logger: deps.Logger,
	})

	userHandler := httpAdapter.NewHandlerUser(&httpAdapter.HandlerUserDeps{
		BaseHandler:  baseHandler,
		IUserUseCase: userUC,
	})

	// Initialize Routers
	activeHandlers := &ActiveHandlers{
		HandlerUser: userHandler,
	}

	activeMiddlewares := &ActiveMiddlewares{}

	InitRouters(router, activeHandlers, activeMiddlewares)

	return &HttpServer{
		conf:   deps.HTTPServer,
		log:    deps.Logger,
		router: router,
	}
}

func (s *HttpServer) Start() error {
	log := s.log.With(slog.String("Address", s.conf.Address+":"+s.conf.Port))
	log.Debug("HTTP server: started creating")

	LoggerHTTP := middleware.NewMiddlewareLogging(&middleware.MiddlewareLoggingDeps{
		Logger: s.log,
	})

	s.server = &http.Server{
		Addr: s.conf.Address + ":" + s.conf.Port,
		Handler: middleware.ChainMiddleware(
			middleware.Timeout(s.conf.WriteTimeout),
			middleware.CORS,
			LoggerHTTP.HandlersLog(),
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
