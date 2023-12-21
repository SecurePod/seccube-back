package container

import "github.com/docker/docker/client"

const (
	dockerClientVersion = "1.42"
)

func (c *ContainerService) CreateDockerClient() (err error) {
	c.Client, err = client.NewClientWithOpts(
		client.FromEnv,
		client.WithVersion(dockerClientVersion),
	)
	return
}
