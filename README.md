# 事前準備

## hosts
* 以下の行を/etc/hostsに追加してください
```
# for develop
127.0.0.1 api mysql
```

# 起動方法
* 開発用サーバの起動
    + ``` make run-dev ```

# Project Initialize
```
go mod init ginapp
```

ref external packakge
```
go mod tidy
```

# Database Migration

## マイグレーションファイルの作成
* コマンド
```
make migrate-create [migration name]
```
* 実行すると database/migrations に up と down のファイルが作成される
```
make migrate-create create_accounts
/migrations/20241229091950_create_accounts.up.sql
/migrations/20241229091950_create_accounts.down.sql
```
* 生成されたファイルの所有者がrootになる場合(linux環境で発生)はchownコマンドで所有者を変更
```
ls -la database/migrations
sudo chown [your user]:[your group] -R database/migrations
```

## マイグレーション適用
* コマンド
```
make migrate-up
```

# Generate model
* コマンド
```
make generate-model
```
* 実行すると database/models に ファイルが生成される
```
database/models/accounts.gen.go
database/models/schema_migrations.gen.go
```
* 生成されたファイルの所有者がrootになる場合(linux環境で発生)はchownコマンドで所有者を変更
```
ls -la database/migrations
sudo chown [your user]:[your group] -R database/models
```
