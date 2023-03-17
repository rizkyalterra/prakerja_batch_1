package routes

import (
	"prakerja/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("articles", controllers.ArticlesGetController)
	e.POST("articles", controllers.ArticlesCreateController)
}
