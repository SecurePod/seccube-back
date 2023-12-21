package container

import (
	"context"

	"github.com/docker/docker/api/types"
	stop "github.com/docker/docker/api/types/container"
	"github.com/pkg/errors"
)

func (c *ContainerService) DeleteContainer(ctx context.Context, id string) error {
	if err := c.CreateDockerClient(); err != nil {
		return errors.Wrap(err, "create client error")
	}
	defer c.Client.Close()

	err := c.Client.ContainerStop(ctx, id, stop.StopOptions{})
	if err != nil {
		return errors.Wrap(err, "stop container error")
	}

	if c.HostConfig.AutoRemove != true {
		err = c.Client.ContainerRemove(ctx, id, types.ContainerRemoveOptions{})
		if err != nil {
			return errors.Wrap(err, "remove container error")
		}
	}
	return nil
}
