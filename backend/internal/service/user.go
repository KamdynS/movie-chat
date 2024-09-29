package service

import (
	"context"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/kamdyns/movie-chat/internal/model"
	"github.com/kamdyns/movie-chat/internal/repository"
	"github.com/kamdyns/movie-chat/pkg/util"
)

type UserService interface {
	CreateUser(ctx context.Context, req *model.CreateUserReq) (*model.CreateUserRes, error)
	Login(ctx context.Context, req *model.LoginUserReq) (*model.LoginUserRes, error)
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

func (s *userService) CreateUser(ctx context.Context, req *model.CreateUserReq) (*model.CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.userRepo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &model.CreateUserRes{
		ID:       strconv.FormatInt(r.ID, 10),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil
}

func (s *userService) Login(ctx context.Context, req *model.LoginUserReq) (*model.LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	u, err := s.userRepo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": u.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// TODO: Use a proper secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return nil, err
	}

	return &model.LoginUserRes{
		AccessToken: tokenString,
		ID:          strconv.FormatInt(u.ID, 10),
		Username:    u.Username,
	}, nil
}
