package test

import (
	"testing"
)

func TestSetUp(t *testing.T) {
	TestPull(t)
	// cli, err := CreateDockerClient()
	// if err != nil {
	// 	t.Fatalf("Failed to create Docker client: %v", err)
	// 	return
	// }
	// ctx := context.Background()

	// cases := map[string]struct {
	// 	image     string
	// 	expectErr bool
	// }{
	// 	"success": {"httpd:latest", false},
	// 	"fail":    {"", true},
	// }

	// for name, tc := range cases {
	// 	t.Run(name, func(t *testing.T) {
	// 		_, err = cli.ImagePull(ctx, tc.image, types.ImagePullOptions{})
	// 		log.Debug().Str("image", tc.image).Msg("image pulled")
	// 		if (err != nil) != tc.expectErr {
	// 			t.Errorf("ImagePull() error = %v, expectErr %v", err, tc.expectErr)
	// 		}
	// 	})
	// }

	// network_cases := map[string]struct {
	// 	name      string
	// 	expectErr bool
	// }{
	// 	"success": {"test", false},
	// 	"fail":    {"", true},
	// }
	// for name, tc := range network_cases {
	// 	t.Run(name, func(t *testing.T) {
	// 		nid, err := CreateNetwork(ctx, cli, tc.name)
	// 		log.Debug().Str("network", nid).Msg("network created")
	// 		if (err != nil) != tc.expectErr {
	// 			t.Errorf("CreateNetwork error = %v, expectErr %v", err, tc.expectErr)
	// 		}
	// 	})
	// }

	// var id string
	// for _, c := range ContainerList["ssh"] {
	// 	t.Run("create container", func(t *testing.T) {
	// 		c.SetNetworkEndpointConfig(nid)
	// 		containerID, err := c.CreateContainer(ctx, cli)
	// 		if err != nil {
	// 			t.Error(err)
	// 			return
	// 		}
	// 		id = *containerID
	// 		t.Log(id)
	// 	})
	// }
	// log.Debug().Str("id", id).Msg("id")
	// i := NewContainerInformation(id)
	// err = i.SetContainerInformation(ctx, cli)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// jsonBytes, err := json.Marshal(i)
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// log.Info().Msgf("%v", string(jsonBytes))

	// assert.Equal(t, "172.17.0.2", i.ContainerIP)
	// assert.Equal(t, []uint16{22, 80}, i.ContainerPorts)
	// assert.Equal(t, []uint16{2222, 8888}, i.HostPorts)

	// DeleteNetwork(ctx, cli, nid)
	// log.Debug().Str("network", nid).Msg("network deleted")
	// for _, c := range ContainerList["ssh"] {
	// 	t.Run("delete container", func(t *testing.T) {
	// 		err = c.DeleteContainer(ctx, cli, id)
	// 		if err != nil {
	// 			t.Error(err)
	// 			return
	// 		}
	// 	})
	// }
	// log.Debug().Str("container", id).Msg("container deleted")

}
