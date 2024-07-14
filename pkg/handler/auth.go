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

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := h.service.Authorisation.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
