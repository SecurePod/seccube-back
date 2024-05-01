package container

import (
	"bufio"
	"bytes"
	"context"
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
)

func TestWriteFile(t *testing.T) {
	cli, err := CreateDockerClient()
	if err != nil {
		t.Error(err)
		return
	}
	ctx := context.Background()

	id := "bdafc1976f6fa83f06b432c6253a9544707d1d35b1904b37d8e549a446a455b3"
	c := NewCmdExecuter(id, []string{"cat", "./htdocs/index.html"})
	res, err := c.CreateExecResponse(ctx, cli)
	if err != nil {
		t.Error(err)
		return
	}
	defer res.Close()

	info := &ContainerWriteInfo{
		Id:       "bdafc1976f6fa83f06b432c6253a9544707d1d35b1904b37d8e549a446a455b3",
		FilePath: "./htdocs/index.html",
		Content:  `<h1>Hello World</h1>`,
	}
	err = info.WriteToFile(ctx, cli)
	if err != nil {
		t.Error(err)
		return
	}

	buf := make([]byte, 1024)
	result, err := res.Reader.Read(buf)
	if err != nil {
		t.Error(err)
		return
	}
	scanner := bufio.NewScanner(bytes.NewReader(buf[:result]))
	for scanner.Scan() {
		log.Debug().Str("line", scanner.Text()).Msg("line")
		assert.Equal(t, `<h1>Hello World</h1>`, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		t.Error(err)
		return
	}

}
