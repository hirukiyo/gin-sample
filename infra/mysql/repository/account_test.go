package repository_test

import (
	"context"
	"errors"
	"testing"

	"github.com/hirukiyo/gin-sample/domain"
	"github.com/hirukiyo/gin-sample/infra/mysql/model"
	"github.com/hirukiyo/gin-sample/infra/mysql/repository"
	"github.com/hirukiyo/gin-sample/testutil"
	"gorm.io/gorm"
)

func TestAccountGetByID(t *testing.T) {
	ctx := t.Context()
	db, _, _ := testutil.GetTestDB()
	defer db.Rollback()

	// prepare data
	account1 := &model.Account{
		Name:     "test1",
		Email:    "test1@example.jp",
		Password: "pass",
		Status:   1,
	}
	if err := gorm.G[model.Account](db).Create(context.Background(), account1); err != nil {
		t.Fatal(err)
	}

	// prepare cases
	cases := []struct {
		name string
		id   uint64
		want *model.Account
		err  error
	}{
		{
			name: "OK",
			id:   account1.ID,
			want: account1,
			err:  nil,
		},
		{
			name: "NG - Notfound",
			id:   1,
			want: nil,
			err:  domain.ErrNotFound,
		},
	}

	// prepare test target
	repo := repository.NewAccountRepository(db)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			account, err := repo.GetByID(ctx, c.id)
			if c.err == nil {
				if err != nil {
					t.Errorf("エラーを想定していないのにエラー発生: %s", err)
					return
				}
				// このケースのみ次のチェックを実行
			} else {
				if !errors.Is(err, c.err) {
					t.Errorf("想定していたのと異なるエラー発生: %s", err)
					return
				}
				// 想定通りのエラー発生
				return
			}

			if account == nil {
				t.Error("データがnilで返却された")
				return
			}
			if c.want.Name != account.Name {
				t.Errorf("invalid Name. case: %s, actual: %s", c.want.Name, account.Name)
			}
			if c.want.Email != account.Email {
				t.Errorf("invalid Email. case: %s, actual: %s", c.want.Email, account.Email)
			}
			if c.want.Password != account.Password {
				t.Errorf("invalid Password. case: %s, actual: %s", c.want.Password, account.Password)
			}
			if c.want.Status != account.Status {
				t.Errorf("invalid Status. case: %d, actual: %d", c.want.Status, account.Status)
			}
		})
	}
}
