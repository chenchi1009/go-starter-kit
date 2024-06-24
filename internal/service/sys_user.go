package service

import (
	"context"
	"time"

	v1 "github.com/chenchi1009/go-starter-kit/api/v1"
	"github.com/chenchi1009/go-starter-kit/internal/model"
	"github.com/chenchi1009/go-starter-kit/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *v1.RegisterRequest) (err error)
	Login(ctx context.Context, req *v1.LoginRequest) (token string, err error)
	GetProfile(ctx context.Context, username string) (resp *v1.GetProfileResponse, err error)
	UpdateProfile(ctx context.Context, req *v1.UpdateProfileRequest) (err error)
	List(ctx context.Context, req *v1.ListUserRequest) (resp *v1.ListUserResponse, err error)
}

func NewUserService(
	service *Service,
	userRepo repository.UserReposistory,
) UserService {
	return &userService{
		userRepo: userRepo,
		Service:  service,
	}
}

type userService struct {
	*Service
	userRepo repository.UserReposistory
}

func (s *userService) Register(ctx context.Context, req *v1.RegisterRequest) (err error) {
	// check if username exists
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return v1.ErrInternalServerError
	}
	if user != nil {
		return v1.ErrSysUserAlreadyExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user = &model.User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	return s.userRepo.Create(ctx, user)
}

func (s *userService) Login(ctx context.Context, req *v1.LoginRequest) (token string, err error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return "", v1.ErrInternalServerError
	}
	if user == nil {
		return "", v1.ErrSysUserNotFound
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return "", err
	}

	token, err = s.Service.jwt.GenToken(user.Username, time.Now().Add(time.Hour*24*90))
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) GetProfile(ctx context.Context, username string) (resp *v1.GetProfileResponse, err error) {
	user, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return nil, v1.ErrInternalServerError
	}
	if user == nil {
		return nil, v1.ErrSysUserNotFound
	}

	return &v1.GetProfileResponse{
		User: v1.User{
			ID:       user.ID,
			Username: user.Username,
			Rules:    user.Rules,
		}}, nil
}

func (s *userService) UpdateProfile(ctx context.Context, req *v1.UpdateProfileRequest) (err error) {
	user, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil {
		return v1.ErrInternalServerError
	}
	if user == nil {
		return v1.ErrSysUserNotFound
	}

	user.Rules = req.Rules

	return s.userRepo.Update(ctx, user)
}

func (s *userService) List(ctx context.Context, req *v1.ListUserRequest) (resp *v1.ListUserResponse, err error) {
	users, err := s.userRepo.List(ctx, req.Page, req.PageSize)
	if err != nil {
		return nil, v1.ErrInternalServerError
	}

	return &v1.ListUserResponse{
		Total: len(users),
		Users: func() (us []v1.User) {
			for _, user := range users {
				us = append(us, v1.User{
					ID:       user.ID,
					Username: user.Username,
					Rules:    user.Rules,
				})
			}
			return
		}()}, nil
}
