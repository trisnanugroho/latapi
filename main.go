package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"strconv"
)

type News struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type BaseResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func main() {
	e := echo.New()
	e.GET("/news", NewsController)           // endpoint 1
	e.GET("/news/:id", DetailNewsController) // endpoint 2
	e.POST("/news", AddNewsController)       // endpoint 3
	e.Start(":8000")
}

// controller
func AddNewsController(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")

	var data = News{1, title, content}

	return c.JSON(http.StatusOK, BaseResponse{
		Message: "Success",
		Data:    data,
	})
}

func DetailNewsController(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))

	var data = News{id, "A", "B"}

	return c.JSON(http.StatusOK, BaseResponse{
		Message: "Success",
		Data:    data,
	})
}

func NewsController(c echo.Context) error {
	var negara = c.QueryParam("negara")
	var sort = c.QueryParam("sort")

	var data []News
	data = append(data, News{1, negara, "B"})
	data = append(data, News{2, sort, "D"})
	return c.JSON(http.StatusOK, BaseResponse{
		Message: "Success",
		Data:    data,
	})
}
