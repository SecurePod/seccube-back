package container

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/rs/zerolog/log"
)

// func TestInspect(t *testing.T) {

// 	cli, err := CreateDockerClient()
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	ctx := context.Background()
// 	jsonc, err := cli.ContainerInspect(ctx, "6dc9e0a7b3cb890e5ccdd1439c6c41d88b6b5fe5e49f36e02fc3b81cd7350967")

// 	log.Info().Msgf("%v", jsonc.NetworkSettings.IPAddress)
// 	log.Info().Msgf("%v", jsonc.NetworkSettings.Ports)
// 	exposedPorts := extractExposedPorts(jsonc.NetworkSettings)
// 	log.Printf("Exposed ports: %v", exposedPorts)

// 	for _, portBindings := range jsonc.NetworkSettings.Ports {
// 		for _, binding := range portBindings {
// 			log.Info().Msgf("%v", binding.HostPort)
// 		}
// 	}
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// }

// func extractExposedPorts(networkSettings *types.NetworkSettings) []string {
// 	var ports []string
// 	for port := range networkSettings.Ports {
// 		portStr := port.Port()
// 		portNumber := strings.Split(portStr, "/")[0]
// 		ports = append(ports, portNumber)
// 	}
// 	return ports
// }

func TestX(t *testing.T) {
	cli, err := CreateDockerClient()
	if err != nil {
		t.Error(err)
		return
	}
	ctx := context.Background()
	c := NewContainerInformation("f5dab2aa3e9423b406745dbdece50e6c91a74ebe4ec572fca3e6ee639ef485ec")
	err = c.SetContainerInformation(ctx, cli)
	if err != nil {
		t.Error(err)
		return
	}
	log.Info().Msgf("%v", c)

	// jsonに変換
	jsonBytes, err := json.Marshal(c)
	if err != nil {
		t.Error(err)
		return
	}
	log.Info().Msgf("%v", string(jsonBytes))
}
