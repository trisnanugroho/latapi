package routes

import (
	"latapi/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) *echo.Echo {
	e.GET("/news", controllers.NewsController)           // endpoint 1
	e.GET("/news/:id", controllers.DetailNewsController) // endpoint 2
	e.POST("/news", controllers.AddNewsController)       // endpoint 3
	return e
}
