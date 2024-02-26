package container

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

func CreateNetwork(ctx context.Context, cli *client.Client, name string) (nid string, err error) {
	res, err := cli.NetworkCreate(
		ctx,
		name,
		types.NetworkCreate{
			CheckDuplicate: true,
			Driver:         "bridge",
		},
	)
	if err != nil {
		return "", errors.Wrap(err, "create network error")
	}
	ctx2 := context.Background()
	time.AfterFunc(time.Minute*40, func() {
		if err := DeleteContainer(ctx2, cli, name); err != nil {
			fmt.Println(err)
		}
	})
	return res.ID, nil
}

func (c *ContainerService) SetNetworkEndpointConfig(name string) {
	if c.NetworkingConfig == nil {
		c.NetworkingConfig = &network.NetworkingConfig{
			EndpointsConfig: make(map[string]*network.EndpointSettings),
		}
	}

	c.NetworkingConfig.EndpointsConfig = map[string]*network.EndpointSettings{
		name: {},
	}
}

func (c *ContainerService) SetNetworkEndpointConfigWithAlias(name string) {
	if c.NetworkingConfig == nil {
		c.NetworkingConfig = &network.NetworkingConfig{
			EndpointsConfig: make(map[string]*network.EndpointSettings),
		}
	}
	c.NetworkingConfig.EndpointsConfig = map[string]*network.EndpointSettings{
		name: {
			Aliases: []string{"db"},
		},
	}
}

func DeleteNetwork(ctx context.Context, cli *client.Client, nid string) error {
	err := cli.NetworkRemove(ctx, nid)
	if err != nil {
		return errors.Wrap(err, "remove network error")
	}
	return nil
}

func (c *ContainerService) AttachNetwork(nid string) {
	c.NetworkingConfig.EndpointsConfig = map[string]*network.EndpointSettings{
		"NetworkIDConfig": {
			NetworkID: nid,
		},
	}
}
