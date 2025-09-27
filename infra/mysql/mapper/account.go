package mapper

import (
	"github.com/hirukiyo/gin-sample/domain/entity"
	"github.com/hirukiyo/gin-sample/infra/mysql/models"
)

func ToAccountEntity(model *models.Account) *entity.Account {
	return &entity.Account{
		ID:        model.ID,
		Name:      model.Name,
		Email:     model.Email,
		Password:  model.Password,
		CreatedAt: model.CreatedAt,
		UpdatedAt: model.UpdatedAt,
	}
}

func ToAccountEntities(models []*models.Account) []*entity.Account {
	accountEntities := make([]*entity.Account, len(models))
	for i, accountModel := range models {
		accountEntities[i] = ToAccountEntity(accountModel)
	}
	return accountEntities
}

func ToAccountModel(entity *entity.Account) *models.Account {
	return &models.Account{
		ID:        entity.ID,
		Name:      entity.Name,
		Email:     entity.Email,
		Password:  entity.Password,
		CreatedAt: entity.CreatedAt,
		UpdatedAt: entity.UpdatedAt,
	}
}

func ToAccountModels(entities []*entity.Account) []*models.Account {
	accountModels := make([]*models.Account, len(entities))
	for i, accountEntity := range entities {
		accountModels[i] = ToAccountModel(accountEntity)
	}
	return accountModels
}
