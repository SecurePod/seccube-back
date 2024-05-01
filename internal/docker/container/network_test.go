package container

import (
	"context"
	"testing"
)

func TestNetworkCreate(t *testing.T) {
	cli, err := CreateDockerClient()
	if err != nil {
		t.Error(err)
		return
	}
	ctx := context.Background()

	name := "test"

	id, err := CreateNetwork(ctx, cli, name)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(id)

	err = DeleteNetwork(ctx, cli, id)
	if err != nil {
		t.Error(err)
	}
}

func TestInvalidNetworkName(t *testing.T) {
	cli, err := CreateDockerClient()
	if err != nil {
		t.Error(err)
		return
	}
	ctx := context.Background()

	name := ""

	_, err = CreateNetwork(ctx, cli, name)
	if err == nil {
		t.Error(err)
		return
	}
}
