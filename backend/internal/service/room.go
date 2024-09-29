package service

import (
	"context"
	"time"

	"github.com/kamdyns/movie-chat/internal/model"
	"github.com/kamdyns/movie-chat/internal/repository"
	"github.com/kamdyns/movie-chat/pkg/util"
)

type RoomService interface {
	CreateRoom(ctx context.Context, name string) (*model.Room, error)
	GetRooms(ctx context.Context) ([]*model.Room, error)
	// New methods
	GetRoom(ctx context.Context, id string) (*model.Room, error)
	UpdateRoom(ctx context.Context, room *model.Room) (*model.Room, error)
	DeleteRoom(ctx context.Context, id string) error
	AddMember(ctx context.Context, roomID string, userID int64) error
	RemoveMember(ctx context.Context, roomID string, userID int64) error
	GetRoomMembers(ctx context.Context, roomID string) ([]*model.RoomMember, error)
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

func (s *roomService) CreateRoom(ctx context.Context, name string) (*model.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	roomID, err := util.GenerateRoomID()
	if err != nil {
		return nil, err
	}

	room := &model.Room{
		ID:   roomID,
		Name: name,
	}

	return s.roomRepo.CreateRoom(ctx, room)
}

func (s *roomService) GetRooms(ctx context.Context) ([]*model.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.roomRepo.GetRooms(ctx)
}

// New methods for roomService
func (s *roomService) GetRoom(ctx context.Context, id string) (*model.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.roomRepo.GetRoom(ctx, id)
}

func (s *roomService) UpdateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.roomRepo.UpdateRoom(ctx, room)
}

func (s *roomService) DeleteRoom(ctx context.Context, id string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.roomRepo.DeleteRoom(ctx, id)
}

func (s *roomService) AddMember(ctx context.Context, roomID string, userID int64) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.roomRepo.AddMember(ctx, roomID, userID)
}

func (s *roomService) RemoveMember(ctx context.Context, roomID string, userID int64) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.roomRepo.RemoveMember(ctx, roomID, userID)
}

func (s *roomService) GetRoomMembers(ctx context.Context, roomID string) ([]*model.RoomMember, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.roomRepo.GetRoomMembers(ctx, roomID)
}
