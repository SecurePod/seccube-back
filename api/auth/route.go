package auth

import (
	"github.com/labstack/echo/v4"
	"github.com/malsuke/seccube-back/api/auth/gen"
	"github.com/malsuke/seccube-back/api/auth/register"
)

type AuthApi struct{}

func (a AuthApi) Login(c echo.Context) error {
	return nil
}

func (a AuthApi) Register(c echo.Context) error {
	return register.Register(c)
}

func (a AuthApi) GetSession(c echo.Context) error {
	return nil
}

var _ gen.ServerInterface = AuthApi{}

func InitRoute(e *echo.Echo) {
	d := e.Group("api/v1/auth")

	gen.RegisterHandlers(d, AuthApi{})

}
