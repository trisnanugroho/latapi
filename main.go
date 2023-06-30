package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Hello struct {
	Name string
}

func main() {
	e := echo.New()

	// router
	e.GET("/hello", HelloController)
	e.Start(":8000")
}

// controller
func HelloController(c echo.Context) error {
	var hello = Hello{"TrisnaN"}
	return c.JSON(http.StatusOK, hello)
}
