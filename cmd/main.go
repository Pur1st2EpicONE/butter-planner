package main

import (
	"log/slog"
	"os"

	"github.com/Pur1st2EpicONE/butter-planner/pkg/server"
)

func main() {
	logger := initLogger()
	slog.SetDefault(logger)

	server, err := server.ServerPrep("8080")
	if err != nil {
		slog.Error("Couldn't prepare the server", slog.String("err", err.Error()))
		os.Exit(1)
	}

	if err := server.Run(); err != nil {
		slog.Error("Couldn't start the server", slog.String("err", err.Error()))
		os.Exit(1)
	}
}

func initLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}
