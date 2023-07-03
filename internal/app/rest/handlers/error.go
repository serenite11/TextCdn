package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Error struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Println(message)
	c.AbortWithStatusJSON(statusCode, Error{Message: message, StatusCode: statusCode})
}
