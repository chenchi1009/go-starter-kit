package repository

import (
	"context"

	"github.com/chenchi1009/go-starter-kit/internal/model"
)

type UserReposistory interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	List(ctx context.Context, page, pageSize int) ([]*model.User, error)
}

func NewUserRepository(r *Repository) UserReposistory {
	return &userRepository{Repository: r}
}

type userRepository struct {
	*Repository
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.Transaction(ctx, func(ctx context.Context) error {
		return r.db.WithContext(ctx).Delete(&model.User{}, id).Error
	})
}

func (r *userRepository) GetByID(ctx context.Context, id uint) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Preload("Rules").First(&user, id).Error
	return &user, err
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	err := r.db.WithContext(ctx).Preload("Rules").Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *userRepository) List(ctx context.Context, page, pageSize int) ([]*model.User, error) {
	var users []*model.User
	err := r.db.WithContext(ctx).Preload("Rules").Limit(pageSize).Offset((page - 1) * pageSize).Find(&users).Error
	return users, err
}
