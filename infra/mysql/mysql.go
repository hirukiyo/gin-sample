package mysql

import (
	"fmt"
	"log"
	"os"
	"time"

	mysqldriver "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Connection struct {
	*gorm.DB
}

func NewConnection(
	MysqlUser string,
	MysqlPassword string,
	MysqlHost string,
	MysqlPort int,
	MysqlDatabase string,
	MysqlLogLevel int,
	MysqlMaxIdleConns int,
	MysqlMaxOpenConns int,
	MysqlConnectionMaxLifetime int,
) (*Connection, error) {

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MysqlUser,
		MysqlPassword,
		MysqlHost,
		MysqlPort,
		MysqlDatabase,
	)

	db, err := gorm.Open(mysqldriver.Open(dsn), &gorm.Config{
		Logger: NewLogger(MysqlLogLevel),
	})
	if err != nil {
		return nil, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(MysqlMaxIdleConns)
	sqlDb.SetMaxOpenConns(MysqlMaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(MysqlConnectionMaxLifetime) * time.Second)

	return &Connection{
		DB: db,
	}, nil
}

func NewLogger(logLevel int) logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,               // Slow SQL threshold
			LogLevel:                  logger.LogLevel(logLevel), // Log level
			IgnoreRecordNotFoundError: true,                      // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,                      // Don't include params in the SQL log
			Colorful:                  false,                     // Disable color
		},
	)
}
