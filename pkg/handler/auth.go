package handler

import (
	"log/slog"
	"net/http"

	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var userInfo butterplanner.User

	if err := c.BindJSON(&userInfo); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.ServiceProvider.CreateUser(userInfo)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	slog.Error(message)
	c.AbortWithStatusJSON(statusCode, gin.H{"error": message})
}
