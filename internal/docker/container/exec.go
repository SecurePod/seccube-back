package container

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type CmdExecuter struct {
	Id  string   `json:"id"`
	Cmd []string `json:"cmd"`
}

func NewCmdExecuter(id string, cmd []string) *CmdExecuter {
	return &CmdExecuter{
		Id:  id,
		Cmd: cmd,
	}
}

func (c *CmdExecuter) CreateExecResponse(ctx context.Context, cli *client.Client) (res types.HijackedResponse, err error) {

	config := types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Cmd:          c.Cmd,
	}

	execId, err := cli.ContainerExecCreate(ctx, c.Id, config)
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
	log.Debug().Str("container", c.Id).Msg("exec attached")

	return res, nil
}
