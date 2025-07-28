package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Pur1st2EpicONE/butter-planner/pkg/repository"
	"github.com/Pur1st2EpicONE/butter-planner/pkg/server"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	db, err := initProject()
	if err != nil {
		logFatal("project init failed", err)
	}
	srv := server.InitServer(viper.GetString("port"), db)
	go func() {
		err := srv.Run()
		if err != nil && err != http.ErrServerClosed {
			logFatal("server run failed", err)
		}
	}()
	osSignCh := make(chan os.Signal, 1)
	signal.Notify(osSignCh, syscall.SIGINT, syscall.SIGTERM)
	<-osSignCh

	if err := srv.Shutdown(context.Background()); err != nil {
		logFatal("server shutdown fail", err)
	}

	if err := db.Close(); err != nil {
		logFatal("db connection failed to close properly", err)
	}
}

func initProject() (*sqlx.DB, error) {

	logger := initLogger()
	slog.SetDefault(logger)

	if err := initConfigs(); err != nil {
		return nil, err
	}

	db, err := repository.ConnectPostgres(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func initConfigs() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func logFatal(msg string, err error) {
	slog.Error(msg, slog.String("err", err.Error()))
	os.Exit(1)
}
