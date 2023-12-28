package handler

import (
	"docker-api/api/docker/container"
	"log"

	"github.com/labstack/echo/v4"
)

func Inspect(c echo.Context) error {
	var i []*container.ContainerInformation
	if err := c.Bind(&i); err != nil {
		return err
	}

	cli, err := container.CreateDockerClient()
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	for _, v := range i {
		v.SetContainerInformation(ctx, cli)
		log.Printf("%v", v)
	}

	return c.JSON(200, i)
}
