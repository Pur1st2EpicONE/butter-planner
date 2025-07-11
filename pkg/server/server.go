package server

import (
	"net/http"
	"time"

	"github.com/Pur1st2EpicONE/butter-planner/pkg/handler"
	"github.com/Pur1st2EpicONE/butter-planner/pkg/repository"
	"github.com/Pur1st2EpicONE/butter-planner/pkg/service"
)

type Server struct {
	httpServer *http.Server
}

func InitServer(port string) *Server {
	repo := repository.NewRepository()
	service := service.NewService(repo)
	handler := handler.NewHandler(service)
	router := handler.InitRoutes()
	server := new(Server)
	server.serverPrep(port, router)
	return server
}

func (s *Server) serverPrep(port string, handler http.Handler) {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}
