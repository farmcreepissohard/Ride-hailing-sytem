package ws

import (
	"context"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/redis/go-redis/v9"
)

type Hub struct {
	mu          sync.RWMutex
	Clients     map[string]*websocket.Conn
	redisClient *redis.Client
}

func NewHub(redisClient *redis.Client) *Hub {
	return &Hub{
		Clients:     make(map[string]*websocket.Conn),
		redisClient: redisClient,
	}
}

func (h *Hub) Add(driverId string, conn *websocket.Conn) {
	h.mu.Lock()
	h.Clients[driverId] = conn
	h.mu.Unlock()

	go h.subscribeToRedis(context.Background(), driverId, conn)
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

func (h *Hub) subscribeToRedis(ctx context.Context, driverId string, conn *websocket.Conn) {
	channelName := "ws_" + driverId
	pubsub := h.redisClient.Subscribe(ctx, channelName)
	defer pubsub.Close()

	for msg := range pubsub.Channel() {
		if err := conn.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
			break
		}
	}
}
