package test

import (
	"context"
	. "docker-api/api/docker/container"
	"docker-api/utils"
	"strconv"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestInspect(t *testing.T) {
	cli, err := CreateDockerClient()
	if err != nil {
		t.Fatalf("Failed to create Docker client: %v", err)
	}
	ctx := context.Background()

	nid := utils.GenerateUUID()

	nid, err = CreateNetwork(ctx, cli, nid)
	if err != nil {
		t.Error(err)
		return
	}
	defer DeleteNetwork(ctx, cli, nid)

	for _, c := range ContainerList["httpd"] {
		c.SetNetworkEndpointConfig(nid)
		id, err := c.CreateContainer(ctx, cli)
		if err != nil {
			t.Error(err)
			return
		}

		con := &ContainerInformation{
			ID: *id,
		}
		err = con.SetContainerInformation(ctx, cli)
		if err != nil {
			t.Error(err)
			return
		}
		assert.NotEmpty(t, con.ContainerIP)
		log.Debug().Str("ip", con.ContainerIP).Msg("container ip")
		for i := 0; i <= len(c.HostConfig.PortBindings)-1; i++ {
			log.Debug().Str("port", strconv.Itoa(int(con.HostPorts[i]))).Msg("container port")
			assert.Equal(t, c.HostConfig.PortBindings["80/tcp"][i].HostPort, strconv.Itoa(int(con.HostPorts[i])))
		}

		err = DeleteContainer(ctx, cli, *id)
		if err != nil {
			t.Error(err)
			return
		}
	}
}
