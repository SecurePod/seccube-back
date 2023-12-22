package container

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	. "github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

var (
	ssh = []*ContainerService{
		NewContainerWithConfig(
			&Config{
				Image: "ubuntu:latest",
				Cmd:   []string{"/bin/bash"},
				Tty:   true,
			},
			&HostConfig{
				PortBindings: nat.PortMap{
					"22/tcp": []nat.PortBinding{
						{
							HostPort: "0",
						},
					},
				},
			},
			nil,
			nil,
		),
		NewContainerWithConfig(
			&Config{
				Image: "ubuntu:latest",
				Cmd:   []string{"/bin/bash"},
				Tty:   true,
			},
			&HostConfig{
				PortBindings: nat.PortMap{
					"22/tcp": []nat.PortBinding{
						{
							HostPort: "0",
						},
					},
				},
			},
			nil,
			nil,
		),
	}
)

func TestCreateMultiple(t *testing.T) {
	ctx := context.Background()

	for _, container := range ssh {
		t.Run("create container", func(t *testing.T) {
			image := container.Config.Image
			_, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
			if err != nil {
				t.Error(err)
			}
			id, err := container.CreateContainer(ctx, cli)
			fmt.Println(id)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
