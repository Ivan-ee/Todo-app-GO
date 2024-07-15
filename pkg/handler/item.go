package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todo"
)

func (h *Handler) createItem(c *gin.Context) {
	userID, err := getUserID(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, "invalid list id param", http.StatusBadRequest)
		return
	}

	var input todo.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, "invalid item param", http.StatusBadRequest)
		return
	}

	id, err := h.service.TodoItem.Create(userID, listId, input)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, "invalid list id param", http.StatusBadRequest)
		return
	}

	items, err := h.service.TodoItem.GetAll(userId, listId)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, "invalid list id param", http.StatusBadRequest)
		return
	}

	item, err := h.service.TodoItem.GetById(userId, itemId)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, "invalid id param", http.StatusBadRequest)
		return
	}

	var input todo.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.service.TodoItem.Update(userId, id, input); err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, StatusResponse{"ok"})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, "invalid list id param", http.StatusBadRequest)
		return
	}

	err = h.service.TodoItem.Delete(userId, itemId)
	if err != nil {
		newErrorResponse(c, err.Error(), http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, StatusResponse{"ok"})
}
