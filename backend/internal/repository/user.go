package repository

import (
	"context"
	"database/sql"

	"github.com/kamdyns/movie-chat/internal/model"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *model.User) (*model.User, error)
	GetUserByClerkID(ctx context.Context, clerkUserID string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	query := `INSERT INTO users(clerk_user_id, username, email) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, user.ClerkUserID, user.Username, user.Email).Scan(&user.ID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByClerkID(ctx context.Context, clerkUserID string) (*model.User, error) {
	user := &model.User{}
	query := `SELECT id, clerk_user_id, username, email FROM users WHERE clerk_user_id = $1`
	err := r.db.QueryRowContext(ctx, query, clerkUserID).Scan(&user.ID, &user.ClerkUserID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, user *model.User) error {
	query := `UPDATE users SET username = $1, email = $2 WHERE clerk_user_id = $3`
	_, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.ClerkUserID)
	return err
}
