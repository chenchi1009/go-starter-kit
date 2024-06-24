package repository

import (
	"context"

	"github.com/chenchi1009/go-starter-kit/internal/model"
)

type RuleRepository interface {
	Create(ctx context.Context, rule *model.Rule) error
	Update(ctx context.Context, rule *model.Rule) error
	Delete(ctx context.Context, id uint) error
	GetByID(ctx context.Context, id uint) (*model.Rule, error)
	GetByName(ctx context.Context, name string) (*model.Rule, error)
	ListAll(ctx context.Context) ([]*model.Rule, error)
}

func NewRuleRepository(r *Repository) RuleRepository {
	return &ruleRepository{Repository: r}
}

type ruleRepository struct {
	*Repository
}

func (r *ruleRepository) Create(ctx context.Context, rule *model.Rule) error {
	return r.db.WithContext(ctx).Create(rule).Error
}

func (r *ruleRepository) Update(ctx context.Context, rule *model.Rule) error {
	return r.db.WithContext(ctx).Save(rule).Error
}

func (r *ruleRepository) Delete(ctx context.Context, id uint) error {
	return r.Transaction(ctx, func(ctx context.Context) error {
		return r.db.WithContext(ctx).Delete(&model.Rule{}, id).Error
	})
}

func (r *ruleRepository) GetByID(ctx context.Context, id uint) (*model.Rule, error) {
	var rule model.Rule
	err := r.db.WithContext(ctx).Preload("Menus").First(&rule, id).Error
	return &rule, err
}

func (r *ruleRepository) GetByName(ctx context.Context, name string) (*model.Rule, error) {
	var rule model.Rule
	err := r.db.WithContext(ctx).Preload("Menus").Where("name = ?", name).First(&rule).Error
	return &rule, err
}

func (r *ruleRepository) ListAll(ctx context.Context) ([]*model.Rule, error) {
	var rules []*model.Rule
	err := r.db.WithContext(ctx).Preload("Menus").Find(&rules).Error
	return rules, err
}
