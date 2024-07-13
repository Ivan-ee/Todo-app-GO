package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.Authorisation.CreateUser(input)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) signIn(c *gin.Context) {

}
