package repository

import (
	"context"

	"github.com/hirukiyo/gin-sample/domain/entity"
)

type AccountRepository interface {
	GetByID(ctx context.Context, id uint64) (*entity.Account, error)
}
