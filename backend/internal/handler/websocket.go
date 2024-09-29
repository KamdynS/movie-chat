package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kamdyns/movie-chat/internal/model"
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomID := c.Query("roomId")
	clientID := c.Query("userId")
	username := c.Query("username")

	client := &ws.Client{
		Conn:     conn,
		Message:  make(chan *ws.Message, 10),
		ID:       clientID,
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Register <- client

	message := &ws.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: username,
	}

	h.hub.Broadcast <- message // Changed from []byte(message.Content) to message

	go client.WriteMessage()
	client.ReadMessage(h.hub)
}

func (h *WebSocketHandler) CreateRoom(c *gin.Context) {
	var req model.CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[req.ID] = &ws.Room{
		ID:      req.ID,
		Name:    req.Name,
		Clients: make(map[string]*ws.Client),
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room created successfully"})
}

func (h *WebSocketHandler) JoinRoom(c *gin.Context) {
	// This is handled in HandleWebSocket, so we can remove this or keep it as a placeholder
	c.JSON(http.StatusOK, gin.H{"message": "Use WebSocket connection to join a room"})
}

func (h *WebSocketHandler) GetRooms(c *gin.Context) {
	rooms := make([]model.RoomRes, 0)
	for _, r := range h.hub.Rooms {
		rooms = append(rooms, model.RoomRes{
			ID:   r.ID,
			Name: r.Name,
		})
	}

	c.JSON(http.StatusOK, rooms)
}

func (h *WebSocketHandler) GetClients(c *gin.Context) {
	var clients []model.ClientRes
	roomID := c.Param("roomId")

	if room, ok := h.hub.Rooms[roomID]; ok {
		for _, cl := range room.Clients {
			clients = append(clients, model.ClientRes{
				ID:       cl.ID,
				Username: cl.Username,
			})
		}
		c.JSON(http.StatusOK, clients)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
	}
}
