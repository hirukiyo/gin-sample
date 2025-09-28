package testutil

import (
	"gorm.io/gorm"

	"github.com/hirukiyo/gin-sample/apiserver/environment"
	"github.com/hirukiyo/gin-sample/infra/mysql"
)

func GetTestEnvironment() (*environment.AppEnvironment, error) {
	return environment.LoadAppEnvironment()
}

func GetTestDB() (*gorm.DB, *environment.AppEnvironment, error) {
	env, err := GetTestEnvironment()
	if err != nil {
		return nil, nil, err
	}
	db, err := mysql.NewConnection(
		env.MysqlUser,
		env.MysqlPassword,
		env.MysqlHost,
		env.MysqlPort,
		env.MysqlDatabase,
		env.MysqlLogLevel,
		env.MysqlMaxIdleConns,
		env.MysqlMaxOpenConns,
		env.MysqlConnectionMaxLifetime,
	)
	if err != nil {
		return nil, nil, err
	}
	return db.Begin(), env, nil
}
