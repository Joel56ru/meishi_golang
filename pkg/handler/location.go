package handler

import (
	"github.com/gin-gonic/gin"
	"meishi_golang/senti"
	"net/http"
)

// @Summary getGeoCode Получение gps координат по адресу
// @Description Получение gps координат по полному адресу
// @Tags Стандартизация, Геометки
// @Accept json
// @Produce json
// @Param text path string true "Начало адреса"
// @Param input body godev.AdressTextSearchInput true "Вводные данные"
// @Success 200 {object} godev.GeoCode
// @Failure 401,500 {object} errorjson
// @Router /dadata/geo [post]
// @Security ApiKeyAuth
func (h *Handler) getGeoCodePost(c *gin.Context) {
	var input senti.AdressTextSearchInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	resp, err := h.services.GetGeoCode(input.Input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, resp)
}
