package repository

import (
	"context"

	"github.com/hirukiyo/gin-sample/domain/entity"
)

type AccountFindConditions struct {
	Name   string
	Email  string
	Status int32
}

type AccountRepository interface {
	Find(ctx context.Context, cond *AccountFindConditions) ([]*entity.Account, error)
	GetByID(ctx context.Context, id uint64) (*entity.Account, error)
	Create(ctx context.Context, account *entity.Account) (uint64, error)
	Update(ctx context.Context, account *entity.Account) (int, error)
	Delete(ctx context.Context, id uint64) (int, error)
}
