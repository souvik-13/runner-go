package main

import (
	"log"
	"sync"
)

type ClientManager struct {
	clients    map[string]*Client
	register   chan *Client
	// reconnect  chan *Client
	unregister chan *Client
	broadcast  chan []byte
	mu         sync.Mutex
}

func NewClientManager() *ClientManager {
	return &ClientManager{
		clients:    make(map[string]*Client),
		register:   make(chan *Client),
		// reconnect:  make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}

func (manager *ClientManager) Start() {
	for {
		select {
		case client := <-manager.register:
			manager.mu.Lock()
			manager.clients[client.GetID()] = client
			manager.mu.Unlock()
			log.Printf("Client registered: %s", client.GetID())

		case client := <-manager.unregister:
			manager.mu.Lock()
			if _, ok := manager.clients[client.GetID()]; ok {
				delete(manager.clients, client.GetID())
				if err := client.Close(); err != nil {
					log.Printf("Error closing client connection: %v", err)
				}
				log.Printf("Client unregistered: %s", client.GetID())
			}
			manager.mu.Unlock()

			// case message := <-manager.broadcast:
			// 	manager.mu.Lock()
			// 	for _, client := range manager.clients {
			// 		select {
			// 		case client.conn.WriteMessage(websocket.TextMessage, message):
			// 		default:
			// 			if _, ok := manager.clients[client.GetID()]; ok {
			// 				delete(manager.clients, client.GetID())
			// 				if err := client.Close(); err != nil {
			// 					log.Printf("Error closing client connection: %v", err)
			// 				}
			// 				log.Printf("Client unregistered: %s", client.GetID())
			// 			}
			// 		}
			// 	}
			// 	manager.mu.Unlock()
		}
	}
}

func (manager *ClientManager) RegisterClient(client *Client) {
	manager.register <- client
}

// func (manager *ClientManager) ReconnectClient(client *Client) {
// 	manager.reconnect <- client
// }

func (manager *ClientManager) UnregisterClient(client *Client) {
	manager.unregister <- client
}

func (manager *ClientManager) BroadcastMessage(message []byte) {
	manager.broadcast <- message
}
