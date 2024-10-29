package main

import (
	"flag"
	"fmt"
	"log"

	"net/http"
)

var (
	port = flag.String("port", "8080", "Port to listen on")
)

func main() {
	flag.Parse()
	var addr = fmt.Sprintf(":%s", *port)
	
	http.HandleFunc("/ws", wsHandler)

	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(addr, nil))
}
