package container

import (
	"context"
	"docker-api/utils"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/rs/zerolog/log"
)

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
				Resources: container.Resources{
					Memory: 512 * 1024 * 1024,
				},
				PortBindings: nat.PortMap{
					"22/tcp": []nat.PortBinding{
						{
							HostPort: "0",
						},
					},
					"2222/tcp": []nat.PortBinding{
						{
							HostPort: "2222",
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

	sqli = []*ContainerService{
		NewContainerWithConfig(
			&container.Config{
				Image: "sqli-app:latest",
			},
			&container.HostConfig{
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
		),
		NewContainerWithConfig(
			&container.Config{
				Image: "sqli-db:latest",
			},
			nil,
			nil,
			nil,
		),
	}

	ssh = []*ContainerService{
		NewContainerWithConfig(
			&container.Config{
				Image: "ssh-ubuntu",
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
				Image: "ssh-ubuntu",
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
							HostPort: "0",
						},
					},
					"2222/tcp": []nat.PortBinding{
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

	ctf = []*ContainerService{
		NewContainerWithConfig(
			&container.Config{
				Image: "build-db",
			},
			&container.HostConfig{
				PortBindings: nat.PortMap{
					"3306/tcp": []nat.PortBinding{
						{
							HostPort: "20002",
						},
					},
				},
			},
			nil,
			nil,
		),
		NewContainerWithConfig(
			&container.Config{
				Image: "build-app",
			},
			&container.HostConfig{
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
		),
	}

	ContainerList = map[string][]*ContainerService{
		"httpd":  httpd,
		"ubuntu": ubuntu,
		"ssh":    ssh,
		"ctf":    ctf,
		"sqli":   sqli,
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
				err = DeleteContainer(ctx, cli, *id)
				if err != nil {
					t.Error(err)
					return
				}
			})
		}
	}
}

func TestCreateWithNetwork(t *testing.T) {
	nid := utils.GenerateUUID()
	var err error
	_, err = CreateNetwork(ctx, cli, nid)
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range ContainerList["sqli"] {
		t.Run("create container", func(t *testing.T) {
			if c.Config.Image == "sqli-db:latest" {
				c.SetNetworkEndpointConfigWithAlias(nid)
			} else {
				c.SetNetworkEndpointConfig(nid)
			}
			id, err := c.CreateContainer(ctx, cli)
			if err != nil {
				t.Error(err)
				return
			}
			t.Log(id)
		})
	}
}
