package handler

import (
	"log/slog"
	"net/http"

	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var userInfo butterplanner.User

	if err := c.ShouldBind(&userInfo); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	_, err := h.service.ServiceProvider.CreateUser(userInfo) // TODO: drop ID return from CreateUser
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}

func (h *Handler) signIn(c *gin.Context) {
	var userInfo butterplanner.LoginPassword
	var userId int

	if err := c.ShouldBind(&userInfo); err != nil {
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

	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/notes")
}

func (h *Handler) authorizeUser(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, "missing token")
		return
	}

	userId, err := h.service.ServiceProvider.ParseToken(token) // TODO: move to a separate function
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set("userId", userId)
}

func (h *Handler) isUserLoggedIn(c *gin.Context) bool {
	token, err := c.Cookie("token")
	if err != nil {
		return false
	}

	_, err = h.service.ServiceProvider.ParseToken(token)
	return err == nil
}

func (h *Handler) logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusSeeOther, "/")
}

func errorResponse(c *gin.Context, statusCode int, message string) {
	slog.Error(message)
	c.AbortWithStatusJSON(statusCode, gin.H{"error": message})
}
