package mapper

import (
	"github.com/hirukiyo/gin-sample/domain/entity"
	"github.com/hirukiyo/gin-sample/infra/mysql/model"
)

func ToAccountEntity(m *model.Account) *entity.Account {
	return &entity.Account{
		ID:        m.ID,
		Name:      m.Name,
		Email:     m.Email,
		Password:  m.Password,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func ToAccountEntities(models []*model.Account) []*entity.Account {
	accountEntities := make([]*entity.Account, len(models))
	for i, accountModel := range models {
		accountEntities[i] = ToAccountEntity(accountModel)
	}
	return accountEntities
}

func ToAccountModel(e *entity.Account) *model.Account {
	return &model.Account{
		ID:        e.ID,
		Name:      e.Name,
		Email:     e.Email,
		Password:  e.Password,
		Status:    e.Status,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func ToAccountModels(entities []*entity.Account) []*model.Account {
	accountModels := make([]*model.Account, len(entities))
	for i, accountEntity := range entities {
		accountModels[i] = ToAccountModel(accountEntity)
	}
	return accountModels
}
