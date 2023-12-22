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

// func NewContainerService(id *string) (*ContainerService, error) {
// 	cli, err := CreateDockerClient()
// 	return &ContainerService{
// 		Id:  *id,
// 		Cli: cli,
// 	}, err
// }

func NewContainerWithConfig(config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, platform *specs.Platform) *ContainerService {
	return &ContainerService{
		Config:           config,
		HostConfig:       hostConfig,
		NetworkingConfig: networkingConfig,
		Platform:         platform,
	}
}

type ContainerInformation struct {
	ID           string
	ExposedPort  int
	InternalPort int
	Ip           string
	Label        string
}

func NewContainerInformation(ID string, ExposedPort int, InternalPort int, Ip string, Label string) *ContainerInformation {
	return &ContainerInformation{
		ID:           ID,
		ExposedPort:  ExposedPort,
		InternalPort: InternalPort,
		Ip:           Ip,
		Label:        Label,
	}
}
