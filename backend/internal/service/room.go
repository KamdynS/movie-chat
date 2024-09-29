package service

import (
	"context"
	"time"

	"github.com/kamdyns/movie-chat/internal/model"
	"github.com/kamdyns/movie-chat/internal/repository"
)

type RoomService interface {
	CreateRoom(ctx context.Context, req *model.CreateRoomReq) (*model.Room, error)
	GetRooms(ctx context.Context) ([]*model.Room, error)
}

type roomService struct {
	roomRepo repository.RoomRepository
	timeout  time.Duration
}

func NewRoomService(roomRepo repository.RoomRepository) RoomService {
	return &roomService{
		roomRepo: roomRepo,
		timeout:  time.Duration(2) * time.Second,
	}
}

func (s *roomService) CreateRoom(ctx context.Context, req *model.CreateRoomReq) (*model.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	room := &model.Room{
		ID:   req.ID,
		Name: req.Name,
	}

	return s.roomRepo.CreateRoom(ctx, room)
}

func (s *roomService) GetRooms(ctx context.Context) ([]*model.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.roomRepo.GetRooms(ctx)
}
