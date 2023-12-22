package container

import (
	"context"
	"testing"
)

func TestNetworkCreate(t *testing.T) {
	cli, err := CreateDockerClient()
	ctx := context.Background()

	name := "test"

	id, err := CreateNetwork(ctx, cli, name)
	if err != nil {
		t.Error(err)
	}
	t.Log(id)

	err = DeleteNetwork(ctx, cli, id)
	if err != nil {
		t.Error(err)
	}
}

func TestInvalidNetworkName(t *testing.T) {
	cli, err := CreateDockerClient()
	ctx := context.Background()

	name := ""

	id, err := CreateNetwork(ctx, cli, name)
	if err == nil {
		t.Error(err)
	}
	t.Log(id)

	err = DeleteNetwork(ctx, cli, id)
	if err == nil {
		t.Error(err)
	}
}
