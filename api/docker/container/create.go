package container

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (c *ContainerService) CreateContainer(ctx context.Context) (*string, error) {
	if err := c.CreateDockerClient(); err != nil {
		return nil, errors.Wrap(err, "create client error")
	}
	defer c.Client.Close()

	create, err := c.Client.ContainerCreate(
		ctx,
		c.Config,
		c.HostConfig,
		c.NetworkingConfig,
		c.Platform,
		"",
	)
	if err != nil {
		return nil, errors.Wrap(err, "create container error")
	}
	log.Debug().Str("container", create.ID).Msg("container created")

	err = c.Client.ContainerStart(ctx, create.ID, types.ContainerStartOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "start container error")
	}
	log.Debug().Str("container", create.ID).Msg("container started")

	return &create.ID, nil

}
