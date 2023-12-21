package handler

import (
	"context"
	"docker-api/api/docker/container"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}
)

func WsHandler(w http.ResponseWriter, r *http.Request) {
	c := container.NewContainerService(nil, nil, nil, nil)
	if err := c.CreateDockerClient(); err != nil {
		log.Println(errors.Wrap(err, "create client error"))
		return
	}
	sub := strings.TrimPrefix(r.URL.Path, "/web-socket/ssh")
	_, id := filepath.Split(sub)
	fmt.Println(id)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, err := c.CreateExecResponse(ctx, id)
	if err != nil {
		log.Println(err)
		return
	}
	defer res.Close()

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			res.Conn.Write(message)
		}
	}()

	buf := make([]byte, 1024)
	for {
		n, err := res.Reader.Read(buf)
		if err != nil {
			log.Println("read:", err)
			break
		}
		err = conn.WriteMessage(websocket.TextMessage, buf[:n])
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
