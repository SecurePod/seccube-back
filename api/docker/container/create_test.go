package container

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

var (
	httpd = NewContainerService(
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
			ctx := context.Background()
			id, err := container.CreateContainer(ctx)
			fmt.Println(id)
			if err != nil {
				t.Error(err)
			}
			err = container.DeleteContainer(ctx, *id)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
