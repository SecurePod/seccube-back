package api

import (
	"docker-api/api/docker"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	port = "8081"
)

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, "ok")
	})

	{
		docker.InitRoute(e)
	}

	e.Logger.Fatal(e.Start(":" + port))
}
