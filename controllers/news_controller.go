package controllers

import (
	"net/http"
	"prakerja/lib/repo"
	"prakerja/models"
	"prakerja/utils"

	"github.com/labstack/echo/v4"
)

func ArticlesCreateController(c echo.Context) error {
	var news models.News
	c.Bind(&news)

	if len(news.Title) == 0 {
		return c.JSON(http.StatusBadRequest, nil)
	}
	err := repo.InsertNews(&news)

	if err != nil {
		return GenerateResponseError(c, err, nil)
	}
	return GenerateResponseSuccess(c, utils.SUCCESS_INSERT_DATA, news)
}

func ArticlesGetController(c echo.Context) error {
	data, err := repo.GetNews()
	if err != nil {
		return GenerateResponseError(c, err, nil)
	}

	return GenerateResponseSuccess(c, utils.SUCCESS_GET_DATA, data)
}
