package container

import (
	"context"

	"github.com/docker/docker/api/types"
	stop "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

func (c *ContainerService) DeleteContainer(ctx context.Context, cli *client.Client, id string) error {

	err := cli.ContainerStop(ctx, id, stop.StopOptions{})
	if err != nil {
		return errors.Wrap(err, "stop container error")
	}

	if !c.HostConfig.AutoRemove {
		err = cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{})
		if err != nil {
			return errors.Wrap(err, "remove container error")
		}
	}
	return nil
}
