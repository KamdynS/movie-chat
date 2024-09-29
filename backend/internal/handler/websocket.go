package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	ws "github.com/kamdyns/movie-chat/internal/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Adjust this for production!
	},
}

type WebSocketHandler struct {
	hub *ws.Hub
}

func NewWebSocketHandler(hub *ws.Hub) *WebSocketHandler {
	return &WebSocketHandler{
		hub: hub,
	}
}

func (h *WebSocketHandler) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := &ws.Client{
		Conn: conn,
		Send: make(chan []byte, 256),
	}

	h.hub.Register <- client

	// Implement ReadPump and WritePump methods
}

func (h *WebSocketHandler) CreateRoom(c *gin.Context) {
	// Implementation depends on your websocket setup
	c.JSON(http.StatusOK, gin.H{"message": "Room created"})
}

func (h *WebSocketHandler) JoinRoom(c *gin.Context) {
	// Implementation depends on your websocket setup
	c.JSON(http.StatusOK, gin.H{"message": "Joined room"})
}

func (h *WebSocketHandler) GetRooms(c *gin.Context) {
	// Implementation depends on your websocket setup
	c.JSON(http.StatusOK, gin.H{"rooms": h.hub.Rooms})
}

func (h *WebSocketHandler) GetClients(c *gin.Context) {
	// Implementation depends on your websocket setup
	roomID := c.Param("roomId")
	if room, ok := h.hub.Rooms[roomID]; ok {
		c.JSON(http.StatusOK, gin.H{"clients": room.Clients})
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
	}
}
