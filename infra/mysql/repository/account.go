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
	"github.com/hirukiyo/gin-sample/infra/mysql/model"
)

type AccountRepositoryImpl struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) repository.AccountRepository {
	return &AccountRepositoryImpl{db: db}
}

func (r *AccountRepositoryImpl) Find(ctx context.Context, cond *repository.AccountFindConditions) ([]*entity.Account, error) {
	db := gorm.G[*model.Account](r.db).Where("1 = 1") // convert gorm.Interface type(result of G) to gorm.ChainInterface type(result of Where)

	if cond.Name != "" {
		db = db.Where("name = ?", cond.Name)
	}
	if cond.Email != "" {
		db = db.Where("email = ?", cond.Email)
	}
	if cond.Status > entity.AccountStatusDefault {
		db = db.Where("status = ?", cond.Status)
	}
	db = db.Order("id asc")

	accountModels, err := db.Find(ctx)
	if err != nil {
		applog.Info(ctx, "select error occurred in AccountRepositoryImpl#Find")
		return []*entity.Account{}, err
	}

	return mapper.ToAccountEntities(accountModels), nil
}

func (r *AccountRepositoryImpl) GetByID(ctx context.Context, id uint64) (*entity.Account, error) {
	account, err := gorm.G[*model.Account](r.db).Where("id = ?", id).First(ctx)
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

func (r *AccountRepositoryImpl) Create(ctx context.Context, account *entity.Account) (uint64, error) {
	accountModel := mapper.ToAccountModel(account)
	if err := gorm.G[model.Account](r.db).Create(ctx, accountModel); err != nil {
		applog.Info(ctx, "insert error occurred in AccountRepositoryImpl#Create")
		return 0, err
	}
	return accountModel.ID, nil
}

func (r *AccountRepositoryImpl) Update(ctx context.Context, account *entity.Account) (int, error) {
	accountModel := mapper.ToAccountModel(account)
	rows, err := gorm.G[model.Account](r.db).Where("id = ?", account.ID).Updates(ctx, *accountModel)
	if err != nil {
		applog.Info(ctx, "update error occurred in AccountRepositoryImpl#Update")
		return 0, err
	}
	return rows, nil
}

func (r *AccountRepositoryImpl) Delete(ctx context.Context, account *entity.Account) (int, error) {
	rows, err := gorm.G[model.Account](r.db).Where("id = ?", account.ID).Delete(ctx)
	if err != nil {
		applog.Info(ctx, "delete error occurred in AccountRepositoryImpl#Delete")
		return 0, err
	}
	return rows, nil
}
