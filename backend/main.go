package main

import (
	"fmt"
	"net/http"

	"github.com/tsuki42/realtime-chat-go-react/pkg/websocket"
)

// define the websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade this connection
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		_, _ = fmt.Fprintf(w, "%+v\n", err)
	}

	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("chat app v0.01")
	setupRoutes()
	_ = http.ListenAndServe(":8080", nil)
}
