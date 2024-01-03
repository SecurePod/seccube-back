package docker

import (
	"docker-api/api/docker/handler"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	d := e.Group("api/v1/docker")

	{
		d.POST("/create/:tag", handler.Create)
		d.POST("/inspect", handler.Inspect)
	}

}
