package container

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

func (c *ContainerService) CreateContainer(ctx context.Context, cli *client.Client) (*string, error) {
	create, err := cli.ContainerCreate(
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

	err = cli.ContainerStart(ctx, create.ID, types.ContainerStartOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "start container error")
	}
	log.Debug().Str("container", create.ID).Msg("container started")

	// コンテナを削除する処理
	ctx2 := context.Background()
	time.AfterFunc(time.Minute*30, func() {
		if err := DeleteContainer(ctx2, cli, create.ID); err != nil {
			fmt.Println(err)
		}
	})

	return &create.ID, nil

}
