package container

import (
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

type ContainerService struct {
	Config           *container.Config
	HostConfig       *container.HostConfig
	NetworkingConfig *network.NetworkingConfig
	Platform         *specs.Platform
}

func NewContainerService(Config *container.Config, HostConfig *container.HostConfig, NetworkingConfig *network.NetworkingConfig, Platform *specs.Platform) *ContainerService {
	return &ContainerService{
		Config:           Config,
		HostConfig:       HostConfig,
		NetworkingConfig: NetworkingConfig,
		Platform:         Platform,
	}
}
