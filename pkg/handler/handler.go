package handler

import (
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
	router.GET("/", h.showHomePage)
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.GET("/sign-up", h.showSignUpPage)

		auth.POST("/sign-in", h.signIn)
		auth.GET("/sign-in", h.showSignInPage)

		auth.POST("/logout", h.logout)
	}
	lists := router.Group("/notes", h.authorizeUser)
	{
		lists.GET("/", h.getAllNotes)
		lists.POST("/", h.addNewNote)
	}
	return router
}
