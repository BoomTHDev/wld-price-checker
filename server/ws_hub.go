package server

import (
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type Client struct {
	Conn *websocket.Conn
}

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan float64
	mu         sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan float64),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				client.Conn.Close()
			}
			h.mu.Unlock()

		case price := <-h.broadcast:
			h.mu.Lock()
			for client := range h.clients {
				if err := client.Conn.WriteJSON(map[string]float64{"price": price}); err != nil {
					log.Printf("WebSocket error: %v", err)
					client.Conn.Close()
					delete(h.clients, client)
				}
			}
			h.mu.Unlock()
		}
	}
}

// WebSocket handler
func (s *fiberServer) handleWebSocket(c *websocket.Conn) {
	client := &Client{Conn: c}
	s.hub.register <- client

	// Set up close handler
	defer func() {
		s.hub.unregister <- client
	}()

	// Keep the connection alive
	for {
		if _, _, err := c.ReadMessage(); err != nil {
			break
		}
	}
}
