package repository

import (
	"context"
	"database/sql"

	"github.com/kamdyns/movie-chat/internal/model"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error)
	GetRooms(ctx context.Context) ([]*model.Room, error)
}

type roomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	query := `INSERT INTO rooms(id, name) VALUES ($1, $2) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, room.ID, room.Name).Scan(&room.ID)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *roomRepository) GetRooms(ctx context.Context) ([]*model.Room, error) {
	query := `SELECT id, name FROM rooms`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []*model.Room
	for rows.Next() {
		room := &model.Room{}
		if err := rows.Scan(&room.ID, &room.Name); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}
	return rooms, nil
}
