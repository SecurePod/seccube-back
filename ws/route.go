package ws

import (
	"log"
	"net/http"

	"github.com/malsuke/seccube-back/ws/handler"
)

var (
	addr = ":8080"
)

func Route() {

	http.HandleFunc("/web-socket/ssh/", handler.WsHandler)

	log.Printf("listening on %s...", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
