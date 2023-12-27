package container

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/rs/zerolog/log"
)

type ContainerWriteInfo struct {
	Id       string `json:"id"`
	FilePath string `json:"filePath"`
	Content  string `json:"content"`
}

func (i ContainerWriteInfo) WriteToFile(ctx context.Context, cli *client.Client) error {
	config := types.ExecConfig{
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
		Tty:          false,
		Cmd:          []string{"/bin/bash", "-c", fmt.Sprintf("cat << 'EOF' > %s\n%s\nEOF\n", i.FilePath, i.Content)},
	}

	execId, err := cli.ContainerExecCreate(ctx, i.Id, config)
	if err != nil {
		log.Debug().Msg("Unable to create exec")
		return err
	}

	ExecStartCheck := types.ExecStartCheck{
		Detach: false,
		Tty:    false,
	}

	res, err := cli.ContainerExecAttach(ctx, execId.ID, ExecStartCheck)
	if err != nil {
		log.Debug().Msg("Unable to attach exec")
		return err
	}
	defer res.Close()

	return nil
}
