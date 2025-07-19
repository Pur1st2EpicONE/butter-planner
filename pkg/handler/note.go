package handler

import (
	"errors"
	"net/http"

	butterplanner "github.com/Pur1st2EpicONE/butter-planner"
	"github.com/gin-gonic/gin"
)

type AllNotes struct {
	AllNotes []butterplanner.Note
}

func (h *Handler) addNewNote(c *gin.Context) {
	userId, err := getContextId(c)
	if err != nil {
		errorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var input butterplanner.Note
	if err := c.BindJSON(&input); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateNote(userId, input)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) getAllNotes(c *gin.Context) {
	userId, err := getContextId(c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	notes, err := h.service.GetAllNotes(userId)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, AllNotes{AllNotes: notes})
}

func getContextId(c *gin.Context) (int, error) {
	id, ok := c.Get("userId")
	if !ok {
		return 0, errors.New("user id not found in gin.Context")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
