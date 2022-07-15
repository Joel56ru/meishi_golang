package handler

import (
	"github.com/gin-gonic/gin"
	"meishi_golang/senti"
	"net/http"
)

// @Summary Регистрация пользователя
// @Description Регистрация базового пользователя для тестирования во время разработки
// @Tags Для разработки
// @Accept json
// @Produce json
// @Param input body godev.UserRegister true "Вводные данные"
// @Success 200 {integer} integer ID
// @Failure 400,500 {object} errorjson
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input senti.UserRegister
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
