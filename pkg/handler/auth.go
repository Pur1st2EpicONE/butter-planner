package handler

import (
	"log/slog"
	"net/http"
	"strings"

	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var userInfo butterplanner.User

	if err := c.ShouldBind(&userInfo); err != nil {
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

func (h *Handler) signIn(c *gin.Context) {
	var userInfo butterplanner.LoginPassword
	var userId int

	if err := c.BindJSON(&userInfo); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	userId, err := h.service.ServiceProvider.GetUserId(userInfo)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := h.service.CreateToken(userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) authorizeUser(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if header == "" {
		errorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		errorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		errorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userId, err := h.service.ServiceProvider.ParseToken(headerParts[1])
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", userId)
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	slog.Error(message)
	c.AbortWithStatusJSON(statusCode, gin.H{"error": message})
}
