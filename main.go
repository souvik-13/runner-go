package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"net/http"

	"github.com/gorilla/websocket"
)

var (
	port = flag.String("port", "8080", "Port to listen on")
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections by returning true
		return true
	},
}

func main() {
	flag.Parse()
	var addr = fmt.Sprintf(":%s", *port)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println(err, "Failed to upgrade connection")
			return
		}
		defer conn.Close()

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println(err, "Failed to read message")
				return
			}

			// message parser
			// Unmarshal the received JSON string into the IncomingMessage struct
			var incomingMessage map[string]interface{}
			err = json.Unmarshal(message, &incomingMessage)
			if err != nil {
				log.Println(err, "Failed to unmarshal message")
				return
			}


			err = conn.WriteMessage(messageType, message)
			if err != nil {
				log.Println(err, "Failed to write message")
				return
			}
		}
	})

	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(addr, nil))
}
