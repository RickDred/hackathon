package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/google/martian/v3/log"
)

type errorResponse struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Errorf(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}
