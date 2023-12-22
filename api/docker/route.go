package docker

import (
	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	d := e.Group("api/v1/docker")

	{
		d.POST("/create", func(c echo.Context) error {
			return c.JSON(200, "ok")
		})
	}

}
