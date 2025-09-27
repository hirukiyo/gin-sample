package service

import (
	"context"

	"github.com/hirukiyo/gin-sample/domain/entity"
	"github.com/hirukiyo/gin-sample/domain/repository"
	"gorm.io/gorm"
)

type AccountService interface {
	GetAccountByID(ctx context.Context, id uint64) (*entity.Account, error)
}

type AccountServiceImpl struct {
	db          *gorm.DB
	accountRepo repository.AccountRepository
}

func NewGetAccountService(db *gorm.DB, accountRepo repository.AccountRepository) AccountService {
	return &AccountServiceImpl{
		db:          db,
		accountRepo: accountRepo,
	}
}

func (s *AccountServiceImpl) GetAccountByID(ctx context.Context, id uint64) (*entity.Account, error) {
	return s.accountRepo.GetByID(ctx, id)
}
