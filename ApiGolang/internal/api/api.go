package api

import (
	"log/slog"

	"github.com/EvansTrein/BlockbusterVHS/config"
)

type Api struct {
	conf *config.Config
	log  *slog.Logger
}

type ApiDeps struct {
	conf *config.Config
	log  *slog.Logger
}

func New(deps *ApiDeps) *Api {
	return &Api{
		conf: deps.conf,
		log:  deps.log,
	}
}
