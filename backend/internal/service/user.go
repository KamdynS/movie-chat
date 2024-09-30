package service

import (
	"context"
	"time"

	"github.com/kamdyns/movie-chat/internal/model"
	"github.com/kamdyns/movie-chat/internal/repository"
)

type UserService interface {
	HandleClerkWebhook(ctx context.Context, event *model.ClerkWebhookEvent) error
}

type userService struct {
	userRepo repository.UserRepository
	timeout  time.Duration
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
		timeout:  time.Duration(2) * time.Second,
	}
}

func (s *userService) HandleClerkWebhook(ctx context.Context, event *model.ClerkWebhookEvent) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	switch event.Type {
	case "user.created":
		user := &model.User{
			ClerkUserID: event.Data.ID,
			Username:    event.Data.Username,
			Email:       event.Data.Email,
		}
		_, err := s.userRepo.CreateUser(ctx, user)
		return err
	case "user.updated":
		user := &model.User{
			ClerkUserID: event.Data.ID,
			Username:    event.Data.Username,
			Email:       event.Data.Email,
		}
		return s.userRepo.UpdateUser(ctx, user)
	default:
		return nil // Ignore unhandled event types
	}
}

func (s *userService) GetUserDetails(ctx context.Context, clerkUserID string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	return s.userRepo.GetUserByClerkID(ctx, clerkUserID)
}
