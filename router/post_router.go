package router

import (
	"crud/config"
	"crud/controllers"
	"crud/repository"
	"crud/service"

	"github.com/labstack/echo/v4"
)

func Routes(echo *echo.Echo, conf config.Config) {
	db := config.InitDB(conf)

	repo := repository.NewPostRepository(db)
	service := service.NewPostService(conf, repo)
	controller := controllers.Controller{
		S: service,
	}

	echo.POST("/post", controller.Create)
	echo.GET("/post", controller.FindAll)
	echo.GET("/post/:id", controller.FindOne)
	echo.PUT("/post/:id", controller.Update)
	echo.DELETE("/post/:id", controller.Delete)
}
