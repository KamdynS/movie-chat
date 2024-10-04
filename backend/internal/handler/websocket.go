package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	repository "github.com/kamdyns/movie-chat/internal/repository"
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
	hub            *ws.Hub
	roomService    service.RoomService
	userRepository repository.UserRepository
}

func NewWebSocketHandler(hub *ws.Hub, roomService service.RoomService, userRepository repository.UserRepository) *WebSocketHandler {
	return &WebSocketHandler{
		hub:            hub,
		roomService:    roomService,
		userRepository: userRepository,
	}
}

func (h *WebSocketHandler) HandleWebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomID := c.Query("roomId")
	clerkUserID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Fetch user details from the database
	user, err := h.userRepository.GetUserByClerkID(c.Request.Context(), clerkUserID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user details"})
		return
	}

	client := &ws.Client{
		Conn:     conn,
		Message:  make(chan *ws.Message, 10),
		ID:       user.ClerkUserID,
		RoomID:   roomID,
		Username: user.Username,
	}

	h.hub.Register <- client

	message := &ws.Message{
		Content:  "A new user has joined the room",
		RoomID:   roomID,
		Username: user.Username,
	}

	h.hub.Broadcast <- message

	go client.WriteMessage()
	client.ReadMessage(h.hub)
}

func (h *WebSocketHandler) JoinRoom(c *gin.Context) {
	// This is handled in HandleWebSocket, so we can remove this or keep it as a placeholder
	c.JSON(http.StatusOK, gin.H{"message": "Use WebSocket connection to join a room"})
}

func (h *WebSocketHandler) LeaveRoom(c *gin.Context) {
	// This is handled in HandleWebSocket, so we can remove this or keep it as a placeholder
	c.JSON(http.StatusOK, gin.H{"message": "Use WebSocket connection to leave a room"})
}
