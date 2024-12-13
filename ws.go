package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// 	CheckOrigin: func(r *http.Request) bool {
// 		// Allow all connections by returning true
// 		return true
// 	},
// }

// func wsHandler(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err, "Failed to upgrade connection")
// 		return
// 	}
// 	defer conn.Close()

// 	for {
// 		messageType, message, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err, "Failed to read message")
// 			return
// 		}

// 		// message parser
// 		log.Printf("Received message: %s\n", message)

// 		err = conn.WriteMessage(messageType, message)
// 		if err != nil {
// 			log.Println(err, "Failed to write message")
// 			return
// 		}
// 	}
// }
