package handler

import (
	"github.com/Pur1st2EpicONE/butter-planner/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/test")
	{
		auth.GET("/a", h.signUp)
	}
	return router
}
