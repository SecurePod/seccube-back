package main

import (
	"docker-api/api"
	"docker-api/ws"
)

func main() {
	go api.Run()
	ws.Route()
}
