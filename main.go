package main

import (
	"latapi/configs"
	"latapi/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.InitDB()
	e := echo.New()
	e = routes.InitRoute(e)
	// route
	e.Start(":8000")
}
