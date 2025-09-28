package entity

import (
	"time"
)

// Account mapped from table <accounts>
type Account struct {
	ID        uint64    // アカウントID
	Name      string    // 名前
	Email     string    // メールアドレス
	Password  string    // パスワード
	Status    int32     // ステータス(1: 有効 / それ以外: 無効)
	CreatedAt time.Time // 作成日時
	UpdatedAt time.Time // 更新日時
}

const (
	AccountStatusDefault = iota
	AccountStatusValid
)
