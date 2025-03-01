package api

import (
	"log/slog"

	"github.com/EvansTrein/BlockbusterVHS/config"
	"github.com/EvansTrein/BlockbusterVHS/internal/server"
)

type Api struct {
	conf   *config.Config
	log    *slog.Logger
	server *server.HttpServer
}

type ApiDeps struct {
	*config.Config
	*slog.Logger
}

func New(deps *ApiDeps) *Api {
	httpServer := server.New(&server.HttpServerDeps{
		HTTPServer: &deps.HTTPServer,
		Logger:     deps.Logger,
	})

	return &Api{
		conf:   deps.Config,
		log:    deps.Logger,
		server: httpServer,
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

	a.server = nil
	a.log.Info("api: stop successful")

	return nil
}
