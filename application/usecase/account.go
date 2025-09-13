package usecase

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type AccountInput struct {
	ID       uint64
	Name     string
	Email    string
	Password string
}

type FindAccountsInput struct {
	Keyword   string
	SortKey   string
	SortOrder string
	Offset    int
	Limit     int
}

type AccountOutput struct {
	ID        uint64
	Name      string
	Email     string
	Password  string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type AccountUsecase interface {
	CreateAccount(ctx context.Context, account *AccountInput) (*AccountOutput, error)
	UpdateAccount(ctx context.Context, account *AccountInput) (*AccountOutput, error)
	DeleteAccount(ctx context.Context, id uint64) error
	GetAccount(ctx context.Context, id uint64) (*AccountOutput, error)
	FindAccounts(ctx context.Context, conditions *FindAccountsInput) ([]*AccountOutput, error)
}

type accountUsecase struct {
	db *gorm.DB
}

func NewAccountUsecase(db *gorm.DB) AccountUsecase {
	return &accountUsecase{
		db: db,
	}
}

func (u *accountUsecase) CreateAccount(ctx context.Context, account *AccountInput) (*AccountOutput, error) {
	return nil, nil
}
func (u *accountUsecase) UpdateAccount(ctx context.Context, account *AccountInput) (*AccountOutput, error) {
	return nil, nil
}
func (u *accountUsecase) DeleteAccount(ctx context.Context, id uint64) error {
	return nil
}
func (u *accountUsecase) GetAccount(ctx context.Context, id uint64) (*AccountOutput, error) {
	return nil, nil
}
func (u *accountUsecase) FindAccounts(ctx context.Context, conditions *FindAccountsInput) ([]*AccountOutput, error) {
	return []*AccountOutput{}, nil
}
