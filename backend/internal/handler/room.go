package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kamdyns/movie-chat/internal/model"
	"github.com/kamdyns/movie-chat/internal/service"
)

type RoomHandler struct {
	roomService service.RoomService
}

func NewRoomHandler(roomService service.RoomService) *RoomHandler {
	return &RoomHandler{
		roomService: roomService,
	}
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	var req model.CreateRoomReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	room, err := h.roomService.CreateRoom(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, room)
}

func (h *RoomHandler) GetRooms(c *gin.Context) {
	rooms, err := h.roomService.GetRooms(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rooms)
}
