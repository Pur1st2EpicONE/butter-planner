package handler

import "github.com/gin-gonic/gin"

func (h *Handler) getAllLists(c *gin.Context) {
	id, _ := c.Get("userId")
	c.JSON(200, gin.H{"id": id})
}
