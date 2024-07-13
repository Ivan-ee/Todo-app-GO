package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, message string, statusCode int) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}
