package main

import (
	"log/slog"
	"os"

	"github.com/Pur1st2EpicONE/butter-planner/pkg/server"
	"github.com/spf13/viper"
)

func main() {

	if err := initProject(); err != nil {
		slog.Error("project init failed", slog.String("err", err.Error()))
		os.Exit(1)
	}

	srv := server.InitServer(viper.GetString("port"))

	if err := srv.Run(); err != nil {
		slog.Error("server run failed", slog.String("err", err.Error()))
		os.Exit(1)
	}
}

func initProject() error {
	logger := initLogger()
	slog.SetDefault(logger)
	return initConfigs()
}

func initLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func initConfigs() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
