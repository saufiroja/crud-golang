package main

import (
	"crud/config"
	"crud/router"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	conf := config.Config{}

	router.Routes(e, conf)

	err := e.Start("127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
}
