package model

import (
	"time"

	"github.com/google/uuid"
)

type RoomMember struct {
	ID       int64     `json:"id"`
	RoomID   string    `json:"room_id"`
	UserID   int64     `json:"user_id"`
	JoinedAt time.Time `json:"joined_at"`
}

type Room struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type CreateRoomReq struct {
	Name      string `json:"name"`
	ExpiresIn int64  `json:"expires_in"`
}

type RoomListResponse struct {
	Rooms       []Room `json:"rooms"`
	TotalCount  int    `json:"totalCount"`
	CurrentPage int    `json:"currentPage"`
	TotalPages  int    `json:"totalPages"`
}

type RoomListReq struct {
	Page  int `form:"page,default=1"`
	Limit int `form:"limit,default=20"`
}
