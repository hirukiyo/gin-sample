package repository

import (
	"context"
	"errors"
	"testing"

	"gorm.io/gorm"

	"github.com/hirukiyo/gin-sample/domain"
	"github.com/hirukiyo/gin-sample/domain/entity"
	"github.com/hirukiyo/gin-sample/domain/repository"
	"github.com/hirukiyo/gin-sample/infra/mysql/model"
	"github.com/hirukiyo/gin-sample/testutil"
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
	repo := NewAccountRepository(db)

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

func TestAccountFind(t *testing.T) {
	ctx := t.Context()
	db, _, _ := testutil.GetTestDB()
	defer db.Rollback()

	// cleaning table
	if _, err := gorm.G[model.Account](db).Where("1 = 1").Delete(ctx); err != nil {
		t.Fatal(err)
	}

	// prepare data
	account1 := &model.Account{
		Name:     "account_find_test1",
		Email:    "account_find_test1@example.jp",
		Password: "pass",
		Status:   1,
	}
	account2 := &model.Account{
		Name:     "account_find_test2",
		Email:    "account_find_test2@example.jp",
		Password: "pass",
		Status:   2,
	}
	account3 := &model.Account{
		Name:     "account_find_test1",
		Email:    "account_find_test3@example.jp",
		Password: "pass",
		Status:   2,
	}
	if err := gorm.G[model.Account](db).Create(ctx, account1); err != nil {
		t.Fatal(err)
	}
	if err := gorm.G[model.Account](db).Create(ctx, account2); err != nil {
		t.Fatal(err)
	}
	if err := gorm.G[model.Account](db).Create(ctx, account3); err != nil {
		t.Fatal(err)
	}

	// prepare cases
	cases := []struct {
		name  string
		cond  *repository.AccountFindConditions
		wants []*model.Account
		err   error
	}{
		{
			name:  "OK - by Name",
			cond:  &repository.AccountFindConditions{Name: "account_find_test2"},
			wants: []*model.Account{account2},
			err:   nil,
		},
		{
			name:  "OK - by Email",
			cond:  &repository.AccountFindConditions{Email: "account_find_test1@example.jp"},
			wants: []*model.Account{account1},
			err:   nil,
		},
		{
			name:  "OK - by Status",
			cond:  &repository.AccountFindConditions{Status: 1},
			wants: []*model.Account{account1},
			err:   nil,
		},
		{
			name:  "OK - by Name and Status",
			cond:  &repository.AccountFindConditions{Name: "account_find_test1", Status: 2},
			wants: []*model.Account{account3},
			err:   nil,
		},
		{
			name:  "OK - Notfound",
			cond:  &repository.AccountFindConditions{Name: "not-exist"},
			wants: []*model.Account{},
			err:   nil,
		},
		{
			name:  "OK - no condition",
			cond:  &repository.AccountFindConditions{},
			wants: []*model.Account{account1, account2, account3},
			err:   nil,
		},
	}

	// prepare test target
	repo := NewAccountRepository(db)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			accounts, err := repo.Find(ctx, c.cond)
			if err != nil {
				t.Errorf("エラーを想定していないのにエラー発生: %s", err)
				return
			}
			if len(accounts) != len(c.wants) {
				t.Errorf("件数が想定と異なる. case: %d, actual: %d", len(c.wants), len(accounts))
				return
			}

			for i, want := range c.wants {
				actual := accounts[i]
				if want.Name != actual.Name {
					t.Errorf("invalid Name. case: %s, actual: %s", want.Name, actual.Name)
				}
				if want.Email != actual.Email {
					t.Errorf("invalid Email. case: %s, actual: %s", want.Email, actual.Email)
				}
				if want.Password != actual.Password {
					t.Errorf("invalid Password. case: %s, actual: %s", want.Password, actual.Password)
				}
				if want.Status != actual.Status {
					t.Errorf("invalid Status. case: %d, actual: %d", want.Status, actual.Status)
				}
			}
		})
	}
}

func TestAccountCreate(t *testing.T) {
	ctx := t.Context()
	db, _, _ := testutil.GetTestDB()
	defer db.Rollback()

	repo := NewAccountRepository(db)

	cases := []struct {
		name    string
		account *entity.Account
		wantErr bool
	}{
		{
			name: "OK",
			account: &entity.Account{
				Name:     "test_create",
				Email:    "test_create@example.com",
				Password: "password",
				Status:   1,
			},
			wantErr: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			id, err := repo.Create(ctx, c.account)
			if (err != nil) != c.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if err == nil {
				if id == 0 {
					t.Error("id should not be 0")
				}
				// check if the record is actually created
				var created model.Account
				created, err := gorm.G[model.Account](db).Where("id = ?", id).First(ctx)
				if err != nil {
					t.Errorf("failed to find created account: %v", err)
				}
				if created.Name != c.account.Name {
					t.Errorf("Name is not correct. want: %s, got: %s", c.account.Name, created.Name)
				}
				if created.Email != c.account.Email {
					t.Errorf("Email is not correct. want: %s, got: %s", c.account.Email, created.Email)
				}
				if created.Status != c.account.Status {
					t.Errorf("Status is not correct. want: %d, got: %d", c.account.Status, created.Status)
				}
			}
		})
	}
}

func TestAccountUpdate(t *testing.T) {
	ctx := t.Context()
	db, _, _ := testutil.GetTestDB()
	defer db.Rollback()

	// prepare data
	account1 := &model.Account{
		Name:     "test_update_before",
		Email:    "test_update_before@example.com",
		Password: "password",
		Status:   1,
	}
	if err := gorm.G[model.Account](db).Create(context.Background(), account1); err != nil {
		t.Fatal(err)
	}

	repo := NewAccountRepository(db)

	cases := []struct {
		name    string
		account *entity.Account
		rows    int
		err     error
	}{
		{
			name: "OK",
			account: &entity.Account{
				ID:       account1.ID,
				Name:     "test_update_after",
				Email:    "test_update_after@example.com",
				Password: "new_password",
				Status:   2,
			},
			rows: 1,
		},
		{
			name: "NG - Not Found",
			account: &entity.Account{
				ID:   99999, // 存在しないID
				Name: "test_update_notfound",
			},
			rows: 0,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			rows, err := repo.Update(ctx, c.account)
			if err != nil {
				t.Errorf("Update() error = %v", err)
				return
			}

			// check if the record is actually updated
			if rows != c.rows {
				t.Errorf("rows are not correct. want: %d, got: %d", c.rows, rows)
				return
			} else if rows == 0 {
				return
			}
			var updated model.Account
			updated, err = gorm.G[model.Account](db).Where("id = ?", c.account.ID).First(ctx)
			if err != nil {
				t.Errorf("failed to find updated account: %v", err)
				return
			}
			if updated.Name != c.account.Name {
				t.Errorf("Name is not correct. want: %s, got: %s", c.account.Name, updated.Name)
			}
			if updated.Email != c.account.Email {
				t.Errorf("Email is not correct. want: %s, got: %s", c.account.Email, updated.Email)
			}
			if updated.Password != c.account.Password {
				t.Errorf("Password is not correct. want: %s, got: %s", c.account.Password, updated.Password)
			}
			if updated.Status != c.account.Status {
				t.Errorf("Status is not correct. want: %d, got: %d", c.account.Status, updated.Status)
			}
		})
	}
}
