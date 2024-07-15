package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorisationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorisationHeader)
	if header == "" {
		newErrorResponse(c, "empty header token", http.StatusUnauthorized)
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, "invalid header token", http.StatusUnauthorized)
		return
	}

	userId, err := h.service.Authorisation.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, "invalid header token", http.StatusUnauthorized)
		return
	}

	c.Set(userCtx, userId)

}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, "user id not found", http.StatusInternalServerError)
		return 0, errors.New("user id not found")
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, "user id is of invalid type", http.StatusInternalServerError)
		return 0, errors.New("user id is of invalid type")
	}

	return idInt, nil
}
