package register

import (
	"github.com/labstack/echo/v4"
	"github.com/malsuke/seccube-back/api/auth/gen"
)

func Register(c echo.Context) error {
	body := gen.RegisterJSONBody{}
	if err := c.Bind(&body); err != nil {
		return err
	}
	return c.JSON(200, body)
}
