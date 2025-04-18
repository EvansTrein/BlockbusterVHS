package api

import (
	"log/slog"

	"github.com/EvansTrein/BlockbusterVHS/config"
	"github.com/EvansTrein/BlockbusterVHS/internal/server"
	"github.com/EvansTrein/BlockbusterVHS/pkg/db/sqlite"
)

type Api struct {
	conf   *config.Config
	log    *slog.Logger
	server *server.HttpServer
	db     *sqlite.SqliteDB
}

type ApiDeps struct {
	*config.Config
	*slog.Logger
}

func New(deps *ApiDeps) *Api {
	db, err := sqlite.New(deps.StoragePath, deps.Logger)
	if err != nil {
		panic(err)
	}

	httpServer := server.New(&server.HttpServerDeps{
		HTTPServer: &deps.HTTPServer,
		Logger:     deps.Logger,
		SqliteDB:   db,
	})

	return &Api{
		conf:   deps.Config,
		log:    deps.Logger,
		server: httpServer,
		db:     db,
	}
}

func (a *Api) MustStart() {
	a.log.Debug("api: started")

	a.log.Info("api: successfully started", "port", a.conf.HTTPServer.Port)
	if err := a.server.Start(); err != nil {
		panic(err)
	}
}

func (a *Api) Stop() error {
	a.log.Debug("api: stop started")

	if err := a.server.Stop(); err != nil {
		a.log.Error("failed to stop HTTP server")
		return err
	}

	if err := a.db.Close(); err != nil {
		a.log.Error("failed to close the database connection")
		return err
	}

	a.server = nil
	a.log.Info("api: stop successful")

	return nil
}
