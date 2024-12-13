package main

import (
	"log"

	"github.com/gorilla/websocket"
)

// Define AccessType type with constants
type AccessType int

const (
    Read AccessType = iota
    Write
    Terminal
)

type Client struct {
    id          string
    // accessToken string
    conn        *websocket.Conn
    pool        *ClientManager

    // accessType AccessType
}

// NewClient creates a new client instance
func NewClient(conn *websocket.Conn, id string, accessToken string, accessType AccessType, pool *ClientManager) *Client {
    return &Client{
        conn:        conn,
        id:          id,
        // accessToken: accessToken,
        // accessType:  accessType,
        pool:        pool,
    }
}

// GetID returns the client's ID
func (c *Client) GetID() string {
    return c.id
}

// // GetAccessToken returns the client's access token
// func (c *Client) GetAccessToken() string {
//     return c.accessToken
// }

// SendMessage sends a message to the client
func (c *Client) SendMessage(messageType int, message []byte) error {
    return c.conn.WriteMessage(messageType, message)
}

// ReadMessage reads a message from the client
func (c *Client) ReadMessage() (int, []byte, error) {
    return c.conn.ReadMessage()
}

// Close closes the client connection
func (c *Client) Close() error {
    return c.conn.Close()
}

// HandleError logs and handles errors
func (c *Client) HandleError(err error) {
    if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
        log.Printf("Unexpected close error: %v", err)
    } else {
        log.Println("Error:", err)
    }
}
