package repository

import (
	"context"
	"database/sql"

	"github.com/kamdyns/movie-chat/internal/model"
)

type RoomRepository interface {
	CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error)
	GetRoom(ctx context.Context, id string) (*model.Room, error)
	GetRooms(ctx context.Context, limit, offset int) ([]model.Room, error)
	GetTotalRoomCount(ctx context.Context) (int, error)
	UpdateRoom(ctx context.Context, room *model.Room) (*model.Room, error)
	DeleteRoom(ctx context.Context, id string) error
	AddMember(ctx context.Context, roomID string, userID int64) error
	RemoveMember(ctx context.Context, roomID string, userID int64) error
	GetRoomMembers(ctx context.Context, roomID string) ([]*model.RoomMember, error)
}

type roomRepository struct {
	db *sql.DB
}

func NewRoomRepository(db *sql.DB) RoomRepository {
	return &roomRepository{db: db}
}

func (r *roomRepository) CreateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	query := `INSERT INTO rooms(id, name, created_by, created_at, expires_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, created_by, created_at, expires_at`
	err := r.db.QueryRowContext(ctx, query, room.ID, room.Name, room.CreatedBy, room.CreatedAt, room.ExpiresAt).Scan(&room.ID, &room.Name, &room.CreatedBy, &room.CreatedAt, &room.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *roomRepository) GetRoom(ctx context.Context, id string) (*model.Room, error) {
	query := `SELECT id, name FROM rooms WHERE id = $1`
	var room model.Room
	err := r.db.QueryRowContext(ctx, query, id).Scan(&room.ID, &room.Name)
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomRepository) GetRooms(ctx context.Context, limit, offset int) ([]model.Room, error) {
	query := `
		SELECT id, name, created_by, created_at, expires_at 
		FROM rooms 
		WHERE expires_at > NOW() 
		ORDER BY created_at DESC 
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []model.Room
	for rows.Next() {
		var room model.Room
		if err := rows.Scan(&room.ID, &room.Name, &room.CreatedBy, &room.CreatedAt, &room.ExpiresAt); err != nil {
			return nil, err
		}
		rooms = append(rooms, room)
	}

	return rooms, nil
}

func (r *roomRepository) GetTotalRoomCount(ctx context.Context) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM rooms WHERE expires_at > NOW()"
	err := r.db.QueryRowContext(ctx, query).Scan(&count)
	return count, err
}

func (r *roomRepository) UpdateRoom(ctx context.Context, room *model.Room) (*model.Room, error) {
	query := `UPDATE rooms SET name = $2 WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, room.ID, room.Name)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *roomRepository) DeleteRoom(ctx context.Context, id string) error {
	query := `DELETE FROM rooms WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *roomRepository) AddMember(ctx context.Context, roomID string, userID int64) error {
	query := `INSERT INTO room_members(room_id, user_id) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, roomID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *roomRepository) RemoveMember(ctx context.Context, roomID string, userID int64) error {
	query := `DELETE FROM room_members WHERE room_id = $1 AND user_id = $2`
	_, err := r.db.ExecContext(ctx, query, roomID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *roomRepository) GetRoomMembers(ctx context.Context, roomID string) ([]*model.RoomMember, error) {
	query := `SELECT id, room_id, user_id, joined_at FROM room_members WHERE room_id = $1`
	rows, err := r.db.QueryContext(ctx, query, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []*model.RoomMember
	for rows.Next() {
		member := &model.RoomMember{}
		if err := rows.Scan(&member.ID, &member.RoomID, &member.UserID, &member.JoinedAt); err != nil {
			return nil, err
		}
		members = append(members, member)
	}
	return members, nil
}
