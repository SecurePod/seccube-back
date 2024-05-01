package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/malsuke/seccube-back/internal/docker/container"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(c echo.Context) error {
	cli, err := container.CreateDockerClient()
	if err != nil {
		log.Println(err)
		return err
	}
	sub := strings.TrimPrefix(c.Request().URL.Path, "/web-socket/ssh")
	_, id := filepath.Split(sub)
	fmt.Println(id)

	conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	con := container.NewCmdExecuter(id, []string{"/bin/bash"})

	res, err := con.CreateExecResponse(ctx, cli)
	if err != nil {
		log.Println(err)
		return err
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
	return nil
}
