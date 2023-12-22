package container

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

var (
	ctx    = context.Background()
	cli, _ = CreateDockerClient()

	httpd = NewContainerWithConfig(
		&container.Config{
			Tty:   true,
			Image: "httpd",
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
	)

	ContainerList = map[string]ContainerService{
		"httpd": *httpd,
	}
)

func TestCreate(t *testing.T) {

	for _, container := range ContainerList {
		t.Run("create container", func(t *testing.T) {
			_, err := cli.ImagePull(ctx, "docker.io/library/httpd", types.ImagePullOptions{})
			if err != nil {
				t.Error(err)
			}
			id, err := container.CreateContainer(ctx, cli)
			fmt.Println(id)
			if err != nil {
				t.Error(err)
			}
			err = container.DeleteContainer(ctx, cli, *id)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
