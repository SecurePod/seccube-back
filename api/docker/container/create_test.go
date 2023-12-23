package container

import (
	"context"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/rs/zerolog/log"
)

var (
	ctx    = context.Background()
	cli, _ = CreateDockerClient()

	httpd = []*ContainerService{
		NewContainerWithConfig(
			&container.Config{
				Tty:   true,
				Image: "ubuntu:latest",
			},
			&container.HostConfig{
				AutoRemove: true,
				PortBindings: nat.PortMap{
					"80/tcp": []nat.PortBinding{
						{
							HostPort: "0",
						},
					},
				},
			},
			// &network.NetworkingConfig{
			// 	EndpointsConfig: map[string]*network.EndpointSettings{
			// 		"NetworkIDConfig": {
			// 			NetworkID: "NetworkID",
			// 		},
			// 	},
			// },
			nil,
			nil,
		),
	}

	ContainerList = map[string][]*ContainerService{
		"httpd":  httpd,
		"ubuntu": ssh,
	}
)

func TestPull(t *testing.T) {
	for _, container := range ContainerList {
		for _, c := range container {
			t.Run("pull image", func(t *testing.T) {
				_, err := cli.ImagePull(ctx, c.Config.Image, types.ImagePullOptions{})
				log.Debug().Str("image", c.Config.Image).Msg("image pulled")
				if err != nil {
					t.Error(err)
					return
				}
			})
		}
	}
}

func TestCreate(t *testing.T) {
	for _, container := range ContainerList {
		for _, c := range container {
			t.Run("create container", func(t *testing.T) {
				id, err := c.CreateContainer(ctx, cli)
				if err != nil {
					t.Error(err)
					return
				}
				t.Log(id)
				err = c.DeleteContainer(ctx, cli, *id)
				if err != nil {
					t.Error(err)
					return
				}
			})
		}

	}
}
