package test

import (
	"context"
	"encoding/json"
	"testing"

	. "docker-api/api/docker/container"

	"github.com/docker/docker/api/types"
	"github.com/rs/zerolog/log"
)

func TestSetUp(t *testing.T) {
	cli, err := CreateDockerClient()
	if err != nil {
		t.Fatalf("Failed to create Docker client: %v", err)
		return
	}
	ctx := context.Background()

	testCases := []pullTestCase{
		{"httpd", []string{"httpd:latest"}},
		{"ubuntu", []string{"ubuntu:latest"}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			for _, image := range tc.images {
				_, err := cli.ImagePull(ctx, image, types.ImagePullOptions{})
				log.Debug().Str("image", image).Msg("image pulled")
				if err != nil {
					t.Error(err)
				}
			}
		})
	}

	nid, err := CreateNetwork(ctx, cli, "test")
	if err != nil {
		t.Error(err)
		return
	}
	log.Debug().Str("network", nid).Msg("network created")

	var id string
	for _, c := range ContainerList["test"] {
		t.Run("create container", func(t *testing.T) {
			c.SetNetworkEndpointConfig(nid)
			containerID, err := c.CreateContainer(ctx, cli)
			if err != nil {
				t.Error(err)
				return
			}
			id = *containerID
			t.Log(id)
		})
	}
	log.Debug().Str("id", id).Msg("id")
	i := NewContainerInformation(id)
	err = i.SetContainerInformation(ctx, cli)
	if err != nil {
		t.Error(err)
		return
	}
	jsonBytes, err := json.Marshal(i)
	if err != nil {
		t.Error(err)
		return
	}
	log.Info().Msgf("%v", string(jsonBytes))

	DeleteNetwork(ctx, cli, nid)
	log.Debug().Str("network", nid).Msg("network deleted")
	for _, c := range ContainerList["test"] {
		t.Run("delete container", func(t *testing.T) {
			err = c.DeleteContainer(ctx, cli, id)
			if err != nil {
				t.Error(err)
				return
			}
		})
	}
	log.Debug().Str("container", id).Msg("container deleted")
}
