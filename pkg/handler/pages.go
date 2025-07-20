package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) showHomePage(c *gin.Context) {
	logged := h.isUserLoggedIn(c)
	if logged {
		c.HTML(http.StatusOK, "layout2.html", nil)
	} else {
		c.HTML(http.StatusOK, "layout.html", nil)
	}
}

func (h *Handler) showSignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func (h *Handler) showSignInPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signin.html", nil)
}
