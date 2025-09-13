package usecases

import (
	"context"
	"fmt"
	"time"

	"github.com/hirukiyo/gin-sample/infra/mysql/models"
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
	ID        uint64     `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
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
	account, err := gorm.G[models.Account](u.db).Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}
	return &AccountOutput{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}, nil
}
func (u *accountUsecase) FindAccounts(ctx context.Context, conditions *FindAccountsInput) ([]*AccountOutput, error) {
	return []*AccountOutput{}, nil
}
