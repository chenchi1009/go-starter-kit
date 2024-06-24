package repository

import (
	"context"

	"github.com/chenchi1009/go-starter-kit/internal/model"
)

type MenuRepository interface {
	Create(ctx context.Context, menu *model.Menu) error
	Update(ctx context.Context, menu *model.Menu) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*model.Menu, error)
	List(ctx context.Context, page, pageSize int) ([]*model.Menu, error)
}

func NewMenuRepository(r *Repository) MenuRepository {
	return &menuRepository{Repository: r}
}

type menuRepository struct {
	*Repository
}

func (r *menuRepository) Create(ctx context.Context, menu *model.Menu) error {
	return r.db.WithContext(ctx).Create(menu).Error
}

func (r *menuRepository) Update(ctx context.Context, menu *model.Menu) error {
	return r.db.WithContext(ctx).Save(menu).Error
}

func (r *menuRepository) Delete(ctx context.Context, id uint) error {
	return r.Transaction(ctx, func(ctx context.Context) error {
		return r.db.WithContext(ctx).Delete(&model.Menu{}, id).Error
	})
}

func (r *menuRepository) GetByID(ctx context.Context, id uint) (*model.Menu, error) {
	var menu model.Menu
	err := r.db.WithContext(ctx).First(&menu, id).Error
	return &menu, err
}

func (r *menuRepository) List(ctx context.Context, page, pageSize int) ([]*model.Menu, error) {
	var menus []*model.Menu
	err := r.db.WithContext(ctx).Limit(pageSize).Offset((page - 1) * pageSize).Find(&menus).Error
	return menus, err
}
