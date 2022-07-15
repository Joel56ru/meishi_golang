package handler

import (
	"meishi_golang/senti"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	StaffAccess         = "gettingUser"
)

func (h *Handler) userIdentity(c *gin.Context) {
	if c.ClientIP() == "" {
		newErrorResponse(c, http.StatusBadRequest, "Невалидный заголовок")
		return
	}
	//Получение заголовка авторизации
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "Нет заголовка авторизации")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "Невалидный заголовок")
		return
	}
	if headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "Невалидный заголовок")
		return
	}
	//Парсинг токена
	parsedToken, err := h.services.ParseToken(headerParts[1], c.ClientIP())
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, "Невалидный токен")
		return
	}
	c.Set(StaffAccess, parsedToken)
}

//Преобразование типа
func getUserJWT(c *gin.Context) senti.UserJWT {
	var userAC senti.UserJWT
	user, ok := c.Get(StaffAccess)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "Argument not found")
		return userAC
	}
	userAC = user.(senti.UserJWT)
	return userAC
}
