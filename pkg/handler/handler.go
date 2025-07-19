package handler

import (
	"net/http"

	"github.com/Pur1st2EpicONE/butter-planner/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.LoadHTMLGlob("templates/*")

	auth := router.Group("/auth")
	{
		auth.GET("/sign-up", func(c *gin.Context) {
			c.HTML(http.StatusOK, "signup.html", nil)
		})
		auth.POST("/sign-up", h.signUp)

		auth.POST("/sign-in", h.signIn)
	}
	lists := router.Group("/lists", h.authorizeUser)
	{
		lists.GET("/", h.getAllNotes)
		lists.POST("/", h.addNewNote)
	}
	return router
}
