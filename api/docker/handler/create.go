package handler

import (
	"docker-api/api/docker/container"
	. "docker-api/api/docker/container"

	. "github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/labstack/echo/v4"
)

var (
	ssh = []*ContainerService{
		NewContainerWithConfig(
			&Config{
				Image: "ubuntu:latest",
				Cmd:   []string{"/bin/bash"},
				Tty:   true,
			},
			&HostConfig{
				PortBindings: nat.PortMap{
					"22/tcp": []nat.PortBinding{
						{
							HostPort: "0",
						},
					},
				},
			},
			nil,
			nil,
		),
		NewContainerWithConfig(
			&Config{
				Image: "ubuntu:latest",
				Cmd:   []string{"/bin/bash"},
				Tty:   true,
			},
			&HostConfig{
				PortBindings: nat.PortMap{
					"22/tcp": []nat.PortBinding{
						{
							HostPort: "0",
						},
					},
				},
			},
			nil,
			nil,
		),
	}
	ContainerList = map[string][]*ContainerService{
		"ubuntu": ssh,
	}
)

func Create(c echo.Context) error {
	// tag := c.QueryParam("tag")
	tag := "ubuntu"
	ctx := c.Request().Context()
	cli, err := container.CreateDockerClient()
	if err != nil {
		return err
	}

	for _, container := range ContainerList[tag] {
		id, err := container.CreateContainer(ctx, cli)
		if err != nil {
			return err
		}
		c.Logger().Debug(id)
	}
	return nil
}
