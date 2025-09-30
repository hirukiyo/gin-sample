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

func NewConnection(
	mysqlUser string,
	mysqlPassword string,
	mysqlHost string,
	mysqlPort int,
	mysqlDatabase string,
	mysqlLogLevel int,
	mysqlMaxIdleConns int,
	mysqlMaxOpenConns int,
	mysqlConnectionMaxLifetime int,
) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlUser,
		mysqlPassword,
		mysqlHost,
		mysqlPort,
		mysqlDatabase,
	)

	db, err := gorm.Open(mysqldriver.Open(dsn), &gorm.Config{
		Logger: NewLogger(mysqlLogLevel),
	})
	if err != nil {
		return nil, err
	}
	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDb.SetMaxIdleConns(mysqlMaxIdleConns)
	sqlDb.SetMaxOpenConns(mysqlMaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(mysqlConnectionMaxLifetime) * time.Second)

	return db, nil
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
