package main

import (
	"latapi/configs"
	"latapi/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()
	configs.InitDB()
	e := echo.New()
	e = routes.InitRoute(e)
	// route
	e.Start(":8000")
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Failing load .env file")
	}
}
