package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	c.String(200, "TEST")
}
