package container

import (
	"context"
	"testing"
)

func TestCreateMultiple(t *testing.T) {
	ctx := context.Background()

	for _, container := range ssh {
		t.Run("create container", func(t *testing.T) {
			id, err := container.CreateContainer(ctx, cli)
			if err != nil {
				t.Error(err)
				return
			}
			err = container.DeleteContainer(ctx, cli, *id)
			if err != nil {
				t.Error(err)
				return
			}
		})
	}
}
