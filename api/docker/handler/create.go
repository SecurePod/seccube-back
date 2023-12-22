package handler

import (
	. "docker-api/api/docker/container"

	. "github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

var (
	ssh = []*ContainerService{
		NewContainerWithConfig(
			&Config{
				Image: "ubuntu:latest",
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
