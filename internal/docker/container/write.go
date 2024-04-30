package container

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

type WriteRequest struct {
	Code string `json:"code"`
	Path string `json:"path"`
	Id   string `json:"id"`
}

func Write(ctx context.Context, cli *client.Client, r WriteRequest) error {
	config := types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Cmd:          []string{"/bin/bash", "-c", fmt.Sprintf("cat << 'EOF' > %s\n%s\nEOF\n", r.Path, r.Code)},
	}

	exec, err := cli.ContainerExecCreate(ctx, r.Id, config)
	if err != nil {
		return errors.Wrap(err, "Unable to create exec")
	}

	ExecStartCheck := types.ExecStartCheck{
		Detach: false,
		Tty:    true,
	}

	resp, err := cli.ContainerExecAttach(ctx, exec.ID, ExecStartCheck)
	if err != nil {
		return errors.Wrap(err, "attach exec error")
	}
	defer resp.Close()

	return nil
}
