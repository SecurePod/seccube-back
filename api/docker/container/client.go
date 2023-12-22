package container

import (
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

const (
	dockerClientVersion = "1.42"
)

func CreateDockerClient() (cli *client.Client, err error) {
	cli, err = client.NewClientWithOpts(
		client.FromEnv,
		client.WithVersion(dockerClientVersion),
	)
	if err != nil {
		err = errors.Wrap(err, "create client error")
	}
	return
}
