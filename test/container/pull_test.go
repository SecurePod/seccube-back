package test

import (
	"context"
	"testing"

	. "docker-api/api/docker/container"

	"github.com/docker/docker/api/types"
	"github.com/rs/zerolog/log"
)

func TestPull(t *testing.T) {
	cli, err := CreateDockerClient()
	if err != nil {
		t.Fatalf("Failed to create Docker client: %v", err)
		return
	}
	ctx := context.Background()

	cases := map[string]struct {
		image     string
		expectErr bool
	}{
		"success": {"httpd:latest", false},
		"fail":    {"", true},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			_, err = cli.ImagePull(ctx, tc.image, types.ImagePullOptions{})
			log.Debug().Str("image", tc.image).Msg("image pulled")
			if (err != nil) != tc.expectErr {
				t.Errorf("ImagePull() error = %v, expectErr %v", err, tc.expectErr)
			}
		})
	}
}
