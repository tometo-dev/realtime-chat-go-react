package main

import (
	"fmt"
	"net/http"

	"github.com/tsuki42/realtime-chat-go-react/pkg/websocket"
)

// define the websocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Websocket endpoint hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		_, _ = fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()

}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("distributed chat app v0.01")
	setupRoutes()
	_ = http.ListenAndServe(":8080", nil)
}
