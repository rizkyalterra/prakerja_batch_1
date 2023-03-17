package controllers

import (
	"net/http"
	"prakerja/models"
	"prakerja/utils"

	"github.com/labstack/echo/v4"
)

func GenerateResponseSuccess(c echo.Context, message string, data interface{}) error {
	return c.JSON(http.StatusOK, models.BaseResponse{
		Message: message,
		Data:    data,
	})
}

func GenerateResponseError(c echo.Context, err error, data interface{}) error {
	return c.JSON(utils.GenerateCodeFromError(err), models.BaseResponse{
		Message: err.Error(),
		Data:    data,
	})
}
