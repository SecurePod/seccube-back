package container

import (
	"context"
	"testing"
)

func TestWrite(t *testing.T) {
	cli, err := CreateDockerClient()
	if err != nil {
		t.Error(err)
		return
	}

	r := &WriteRequest{
		Code: "dwadadwadawdwadawdad",
		Path: "./index.php",
		Id:   "812d6b143ff7",
	}

	err = Write(context.Background(), cli, *r)
	if err != nil {
		t.Error(err)
		return
	}
}
