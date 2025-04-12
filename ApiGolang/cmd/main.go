package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/EvansTrein/BlockbusterVHS/config"
	"github.com/EvansTrein/BlockbusterVHS/internal/api"
	"github.com/EvansTrein/BlockbusterVHS/pkg/logs"
)

func main() {
	var conf *config.Config
	var log *slog.Logger

	conf = config.MustLoad()
	log = logs.InitLog(conf.Env)

	serverApi := api.New(&api.ApiDeps{
		Config: conf,
		Logger: log,
	})

	go func() {
		serverApi.MustStart()
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-done
	if err := serverApi.Stop(); err != nil {
		log.Error("an error occurred when stopping the application", "error", err)
		panic(err)
	}
}
