package container

import (
	"context"

	"github.com/docker/docker/api/types"
	stop "github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func DeleteContainer(ctx context.Context, cli *client.Client, id string) error {

	err := cli.ContainerStop(ctx, id, stop.StopOptions{})
	if err != nil {
		return errors.Wrap(err, "stop container error")
	}
	log.Debug().Str("container", id).Msg("container stopped")

	info, err := cli.ContainerInspect(ctx, id)
	if err != nil {
		return errors.Wrap(err, "inspect container error")
	}

	if !info.HostConfig.AutoRemove {
		err = cli.ContainerRemove(ctx, id, types.ContainerRemoveOptions{})
		if err != nil {
			return errors.Wrap(err, "remove container error")
		}
		log.Debug().Str("container", id).Msg("container removed")
	}
	return nil
}
