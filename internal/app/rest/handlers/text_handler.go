package handlers

import (
	"Test-Task/TextCDN/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) GetText(c *gin.Context) {
	var text models.Text
	err := c.BindJSON(&text)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	text, err = h.services.TextViewer.GetText(text.Type, text.Language)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"docs":        text,
		"status code": http.StatusOK,
		"message":     "successfully",
	})
}
