package server

import (
	"fmt"

	"github.com/Pur1st2EpicONE/butter-planner/pkg/handler"
)

func ServerPrep(port string) (*Server, error) {
	handler := new(handler.Handler)
	router := handler.InitRoutes()
	server := new(Server)
	err := server.InitServer(port, router)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize server: %v", err)
	}
	return server, nil
}
