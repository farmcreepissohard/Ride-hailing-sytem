package transport

import (
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	mu      sync.RWMutex
	Clients map[string]*websocket.Conn
}

func NewHub() *Hub {
	return &Hub{
		Clients: make(map[string]*websocket.Conn),
	}
}

func (h *Hub) Add(driverId string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.Clients[driverId] = conn
}

func (h *Hub) Remove(driverId string, conn *websocket.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if conn, ok := h.Clients[driverId]; ok {
		conn.Close()
		delete(h.Clients, driverId)
	}
}

func (h *Hub) Send(driverId string, message []byte) error {
	h.mu.RLock()
	conn, exist := h.Clients[driverId]
	h.mu.RUnlock()

	if !exist {
		return fmt.Errorf("driver %s is offline ", driverId)
	}
	return conn.WriteJSON(message)
}
