package main

import (
	"log/slog"
	"os"

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

	if err := srv.Run(); err != nil {
		logFatal("server run failed", err)
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
