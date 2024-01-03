package handler

import (
	docker "docker-api/api/docker/container"
	"docker-api/utils"
	"log/slog"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/labstack/echo/v4"
)

var (
	ssh = []*docker.ContainerService{
		docker.NewContainerWithConfig(
			&container.Config{
				Image: "ssh-attack:latest",
			},
			&container.HostConfig{
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
		docker.NewContainerWithConfig(
			&container.Config{
				Image: "ssh-def:latest",
			},
			&container.HostConfig{
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
	ContainerList = map[string][]*docker.ContainerService{
		"sshBrute": ssh,
	}
)

func Create(c echo.Context) error {
	tag := c.Param("tag")
	slog.Info(tag)
	ctx := c.Request().Context()
	cli, err := docker.CreateDockerClient()
	if err != nil {
		return err
	}

	var ids []map[string]string

	nid := utils.GenerateUUID()
	_, err = docker.CreateNetwork(ctx, cli, nid)
	if err != nil {
		return err
	}

	for _, container := range ContainerList[tag] {
		container.SetNetworkEndpointConfig(nid)
		id, err := container.CreateContainer(ctx, cli)
		if err != nil {
			return err
		}
		ids = append(ids, map[string]string{
			"id": *id,
		})
		c.Logger().Debug(id)
	}
	return c.JSON(200, ids)
}
