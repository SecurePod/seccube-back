package container

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (c *ContainerService) CreateExecResponse(ctx context.Context, cli *client.Client, id string) (res types.HijackedResponse, err error) {

	config := types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Cmd:          []string{"/bin/bash"},
	}

	execId, err := cli.ContainerExecCreate(ctx, id, config)
	if err != nil {
		return res, errors.Wrap(err, "create exec error")
	}

	ExecStartCheck := types.ExecStartCheck{
		Detach: false,
		Tty:    true,
	}

	res, err = cli.ContainerExecAttach(ctx, execId.ID, ExecStartCheck)
	if err != nil {
		return res, errors.Wrap(err, "exec attach error")
	}
	log.Debug().Str("container", id).Msg("exec attached")

	return res, nil
}
