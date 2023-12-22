package container

import (
	"context"

	"github.com/docker/docker/api/types"
	stop "github.com/docker/docker/api/types/container"
	"github.com/pkg/errors"
)

func (c *ContainerService) DeleteContainer(ctx context.Context, id string) error {
	cli, err := CreateDockerClient()
	if err != nil {
		return errors.Wrap(err, "create client error")
	}
	defer cli.Close()

	err = cli.ContainerStop(ctx, id, stop.StopOptions{})
	if err != nil {
		return errors.Wrap(err, "stop container error")
	}

	if c.HostConfig.AutoRemove != true {
		err = cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{})
		if err != nil {
			return errors.Wrap(err, "remove container error")
		}
	}
	return nil
}
