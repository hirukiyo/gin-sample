package repository

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/hirukiyo/gin-sample/apiserver/applog"
	"github.com/hirukiyo/gin-sample/domain"
	"github.com/hirukiyo/gin-sample/domain/entity"
	"github.com/hirukiyo/gin-sample/domain/repository"
	"github.com/hirukiyo/gin-sample/infra/mysql/mapper"
	"github.com/hirukiyo/gin-sample/infra/mysql/models"
)

type AccountRepositoryImpl struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) repository.AccountRepository {
	return &AccountRepositoryImpl{db: db}
}

func (r *AccountRepositoryImpl) GetByID(ctx context.Context, id uint64) (*entity.Account, error) {
	account, err := gorm.G[*models.Account](r.db).Where("id = ?", id).First(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			applog.Info(ctx, "not found error in AccountRepositoryImpl#GetByID")
			return nil, domain.ErrNotFound
		}
		applog.Info(ctx, "error occurred in AccountRepositoryImpl#GetByID")
		return nil, err
	}
	return mapper.ToAccountEntity(account), nil
}
