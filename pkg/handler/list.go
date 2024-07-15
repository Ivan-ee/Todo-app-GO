package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo"
)

func (h *Handler) createList(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	var input todo.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.service.TodoList.Create(userID, input)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getAllListResponse struct {
	Data []todo.TodoList `json:"data"`
}

func (h *Handler) getAllLists(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	lists, err := h.service.TodoList.GetAll(userID)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, getAllListResponse{
		Data: lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, "invalid id param", http.StatusBadRequest)
		return
	}

	list, err := h.service.TodoList.GetById(userID, id)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
