package usecases

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/hirukiyo/gin-sample/apiserver/applog"
	"github.com/hirukiyo/gin-sample/domain/service"
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
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type AccountUsecase interface {
	CreateAccount(ctx context.Context, account *AccountInput) (*AccountOutput, error)
	UpdateAccount(ctx context.Context, account *AccountInput) (*AccountOutput, error)
	DeleteAccount(ctx context.Context, id uint64) error
	GetAccount(ctx context.Context, id uint64) (*AccountOutput, error)
	FindAccounts(ctx context.Context, conditions *FindAccountsInput) ([]*AccountOutput, error)
}

type accountUsecaseImpl struct {
	db             *gorm.DB
	accountService service.AccountService
}

func NewAccountUsecase(db *gorm.DB, accountService service.AccountService) AccountUsecase {
	return &accountUsecaseImpl{
		db:             db,
		accountService: accountService,
	}
}

func (u *accountUsecaseImpl) CreateAccount(ctx context.Context, account *AccountInput) (*AccountOutput, error) {
	return nil, nil
}
func (u *accountUsecaseImpl) UpdateAccount(ctx context.Context, account *AccountInput) (*AccountOutput, error) {
	return nil, nil
}
func (u *accountUsecaseImpl) DeleteAccount(ctx context.Context, id uint64) error {
	return nil
}
func (u *accountUsecaseImpl) GetAccount(ctx context.Context, id uint64) (*AccountOutput, error) {
	account, err := u.accountService.GetAccountByID(ctx, id)
	if err != nil {
		applog.Info(ctx, "error occurred in accountUsecaseImpl#GetAccount")
		return nil, err
	}
	return &AccountOutput{
		ID:        account.ID,
		Name:      account.Name,
		Email:     account.Email,
		CreatedAt: account.CreatedAt,
		UpdatedAt: account.UpdatedAt,
	}, nil
}
func (u *accountUsecaseImpl) FindAccounts(ctx context.Context, conditions *FindAccountsInput) ([]*AccountOutput, error) {
	return []*AccountOutput{}, nil
}
