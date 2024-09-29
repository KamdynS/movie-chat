package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kamdyns/movie-chat/internal/model"
	"github.com/kamdyns/movie-chat/internal/service"
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
	hub         *ws.Hub
	roomService service.RoomService // Add this line
}

func NewWebSocketHandler(hub *ws.Hub, roomService service.RoomService) *WebSocketHandler {
	return &WebSocketHandler{
		hub:         hub,
		roomService: roomService, // Add this line
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
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Room name is required"})
		return
	}

	room, err := h.roomService.CreateRoom(c.Request.Context(), name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	h.hub.Rooms[room.ID] = &ws.Room{
		ID:      room.ID,
		Name:    room.Name,
		Clients: make(map[string]*ws.Client),
	}

	c.JSON(http.StatusCreated, room)
}

func (h *WebSocketHandler) JoinRoom(c *gin.Context) {
	// This is handled in HandleWebSocket, so we can remove this or keep it as a placeholder
	c.JSON(http.StatusOK, gin.H{"message": "Use WebSocket connection to join a room"})
}

func (h *WebSocketHandler) GetRooms(c *gin.Context) {
	rooms, err := h.roomService.GetRooms(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rooms)
}

func (h *WebSocketHandler) GetRoom(c *gin.Context) {
	id := c.Param("id")
	room, err := h.roomService.GetRoom(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, room)
}

func (h *WebSocketHandler) UpdateRoom(c *gin.Context) {
	var room model.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room.ID = c.Param("id")
	updatedRoom, err := h.roomService.UpdateRoom(c.Request.Context(), &room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update the room in the hub
	if hubRoom, ok := h.hub.Rooms[updatedRoom.ID]; ok {
		hubRoom.Name = updatedRoom.Name
	}

	c.JSON(http.StatusOK, updatedRoom)
}

func (h *WebSocketHandler) DeleteRoom(c *gin.Context) {
	id := c.Param("id")
	err := h.roomService.DeleteRoom(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Remove the room from the hub
	delete(h.hub.Rooms, id)

	c.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully"})
}

func (h *WebSocketHandler) AddMember(c *gin.Context) {
	roomID := c.Param("id")
	var req struct {
		UserID int64 `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.roomService.AddMember(c.Request.Context(), roomID, req.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member added successfully"})
}

func (h *WebSocketHandler) RemoveMember(c *gin.Context) {
	roomID := c.Param("id")
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = h.roomService.RemoveMember(c.Request.Context(), roomID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Remove the client from the hub room if they're connected
	if room, ok := h.hub.Rooms[roomID]; ok {
		delete(room.Clients, strconv.FormatInt(userID, 10))
	}

	c.JSON(http.StatusOK, gin.H{"message": "Member removed successfully"})
}

func (h *WebSocketHandler) GetRoomMembers(c *gin.Context) {
	roomID := c.Param("id")
	members, err := h.roomService.GetRoomMembers(c.Request.Context(), roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
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
