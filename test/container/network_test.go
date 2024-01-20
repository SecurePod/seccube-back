package test

import (
	"context"
	. "docker-api/api/docker/container"
	"testing"

	"github.com/rs/zerolog/log"
)

func TestCreateNetwork(t *testing.T) {
	cli, err := CreateDockerClient()
	if err != nil {
		t.Fatalf("Failed to create Docker client: %v", err)
	}
	ctx := context.Background()

	var tests = map[string]struct {
		name    string
		wantErr bool
	}{
		"success": {"test", false},
		"fail":    {"", true},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			id, err := CreateNetwork(ctx, cli, tt.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateNetwork() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Debug().Str("network", id).Msg("network created")
			if !tt.wantErr {
				defer DeleteNetwork(ctx, cli, id)
			}
			log.Debug().Str("network", id).Msg("network deleted")
		})
	}
}
