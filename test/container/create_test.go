package test

import (
	"context"
	"testing"

	. "docker-api/api/docker/container"
	"docker-api/utils"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
)

var (
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
							HostPort: "8888",
						},
					},
				},
			},
			nil,
			nil,
		),
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
				},
			},
			nil,
			nil,
		),
	}
	ContainerList = map[string][]*ContainerService{
		"httpd": httpd,
		"fail":  nil,
	}
)

func TestCreateContainer(t *testing.T) {
	cli, err := CreateDockerClient()
	if err != nil {
		t.Fatalf("Failed to create Docker client: %v", err)
	}
	ctx := context.Background()

	cases := map[string]struct {
		name      string
		expectErr bool
	}{
		"success": {"httpd", false},
		"fail":    {"fail", true},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			for _, c := range ContainerList[tc.name] {
				id, err := c.CreateContainer(ctx, cli)
				if (err != nil) != tc.expectErr {
					t.Errorf("CreateContainer() error = %v, expectErr %v", err, tc.expectErr)
					return
				}
				if !tc.expectErr {
					DeleteContainer(ctx, cli, *id)
				}
			}
		})
	}
}

func TestCreateContainerWithNetwork(t *testing.T) {
	cli, err := CreateDockerClient()
	if err != nil {
		t.Fatalf("Failed to create Docker client: %v", err)
	}
	ctx := context.Background()

	cases := map[string]struct {
		name      string
		expectErr bool
	}{
		"first":  {"httpd", false},
		"second": {"httpd", false},
		"fail":   {"fail", true},
	}
	uuid := utils.GenerateUUID()

	nid, err := CreateNetwork(ctx, cli, uuid)
	if err != nil {
		t.Error(err)
		return
	}
	defer DeleteNetwork(ctx, cli, nid)

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {

			for _, c := range ContainerList[tc.name] {
				c.SetNetworkEndpointConfig(nid)
				id, err := c.CreateContainer(ctx, cli)
				if (err != nil) != tc.expectErr {
					t.Errorf("CreateContainer() error = %v, expectErr %v", err, tc.expectErr)
					return
				}
				if !tc.expectErr {
					DeleteContainer(ctx, cli, *id)
				}
			}
		})
	}
}
