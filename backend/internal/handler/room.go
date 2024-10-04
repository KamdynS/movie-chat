package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kamdyns/movie-chat/internal/model"
	"github.com/kamdyns/movie-chat/internal/service"
	"github.com/kamdyns/movie-chat/pkg/util"
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

	roomID, err := util.GenerateRoomID() // Implement this function to generate a unique ID
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate room ID"})
		return
	}

	room := &model.Room{
		ID:   roomID,
		Name: req.Name,
	}

	createdRoom, err := h.roomService.CreateRoom(c.Request.Context(), room)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdRoom)
}

func (h *RoomHandler) GetRooms(c *gin.Context) {
	rooms, err := h.roomService.GetRooms(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rooms)
}

func (h *RoomHandler) GetRoom(c *gin.Context) {
	id := c.Param("id")
	room, err := h.roomService.GetRoom(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, room)
}

func (h *RoomHandler) UpdateRoom(c *gin.Context) {
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

	c.JSON(http.StatusOK, updatedRoom)
}

func (h *RoomHandler) DeleteRoom(c *gin.Context) {
	id := c.Param("id")
	err := h.roomService.DeleteRoom(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Room deleted successfully"})
}

func (h *RoomHandler) AddMember(c *gin.Context) {
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

func (h *RoomHandler) RemoveMember(c *gin.Context) {
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

	c.JSON(http.StatusOK, gin.H{"message": "Member removed successfully"})
}

func (h *RoomHandler) GetRoomMembers(c *gin.Context) {
	roomID := c.Param("id")
	members, err := h.roomService.GetRoomMembers(c.Request.Context(), roomID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
}
