package test

import (
	"context"
	"testing"

	. "docker-api/api/docker/container"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/rs/zerolog/log"
)

type pullTestCase struct {
	name   string
	images []string
}

type createTestCase struct {
	name       string
	containers []*ContainerService
}

var (
	ctx    = context.Background()
	cli, _ = CreateDockerClient()

	ubuntu = []*ContainerService{
		NewContainerWithConfig(
			&container.Config{
				Image: "ubuntu:latest",
				Cmd:   []string{"/bin/bash"},
				Tty:   true,
			},
			&container.HostConfig{
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
			&container.Config{
				Image: "ubuntu:latest",
				Cmd:   []string{"/bin/bash"},
				Tty:   true,
			},
			&container.HostConfig{
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

	test = []*ContainerService{
		NewContainerWithConfig(
			&container.Config{
				Image: "ubuntu:latest",
				Cmd:   []string{"/bin/bash"},
				Tty:   true,
			},
			&container.HostConfig{
				PortBindings: nat.PortMap{
					"2222/tcp": []nat.PortBinding{
						{
							HostPort: "2222",
						},
					},
					"3333/tcp": []nat.PortBinding{
						{
							HostPort: "3333",
						},
					},
				},
			},
			nil,
			nil,
		),
	}

	httpd = []*ContainerService{
		NewContainerWithConfig(
			&container.Config{
				Tty:   true,
				Image: "httpd:latest",
			},
			&container.HostConfig{
				AutoRemove: true,
				PortBindings: nat.PortMap{
					"80/tcp": []nat.PortBinding{
						{
							HostPort: "9999",
						},
					},
					"22/tcp": []nat.PortBinding{
						{
							HostPort: "2222",
						},
					},
				},
			},
			nil,
			nil,
		),
	}
	ContainerList = map[string][]*ContainerService{
		"httpd":  httpd,
		"ubuntu": ubuntu,
		"test":   test,
	}
)

func TestCreateContainer(t *testing.T) {
	testCases := []createTestCase{
		{"httpd", httpd},
		{"ubuntu", ubuntu},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, c := range tc.containers {
				id, err := c.CreateContainer(ctx, cli)
				if err != nil {
					t.Error(err)
					continue
				}
				log.Debug().Str("container", *id).Msg("container created")
				t.Log(id)
				err = c.DeleteContainer(ctx, cli, *id)
				if err != nil {
					t.Error(err)
				}
				log.Debug().Str("container", *id).Msg("container deleted")
			}
		})
	}
}
