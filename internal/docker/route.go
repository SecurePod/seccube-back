package docker

import (
	"github.com/malsuke/seccube-back/internal/docker/handler"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	d := e.Group("api/v1/docker")

	{
		d.POST("/create/:tag", handler.Create)
		d.POST("/inspect", handler.Inspect)
		d.POST("/write", handler.Write)
	}

	e.GET("/web-socket/ssh/:id", handler.WsHandler)

}
