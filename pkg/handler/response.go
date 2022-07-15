package handler

import (
	"github.com/gin-gonic/gin"
)

//errorJson структура ошибок
type errorJson struct {
	Message string `json:"message"`
	Detail  string `json:"detail"`
}

//newErrorResponse обработчик ошибок
func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorJson{Message: message})
	return
}
