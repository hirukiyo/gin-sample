package app

import (
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type App struct {
	Env    *AppEnvironment
	Engine *gin.Engine
	GormDB *gorm.DB
}

// Environment
// -----------------------------------------------------------------------------
type AppEnvironment struct {
	AppMode                    string   `envconfig:"APP_MODE"`
	AppName                    string   `envconfig:"APP_NAME"`
	AppHost                    string   `envconfig:"APP_HOST"`
	AppPort                    int      `envconfig:"APP_PORT"`
	AppVersion                 string   `envconfig:"APP_VERSION"`
	AppLogLevel                string   `envconfig:"APP_LOG_LEVEL"`
	AppTest                    string   `envconfig:"APP_TEST"`
	MysqlUser                  string   `envconfig:"MYSQL_USER"`
	MysqlPassword              string   `envconfig:"MYSQL_PASSWORD"`
	MysqlHost                  string   `envconfig:"MYSQL_HOST"`
	MysqlPort                  int      `envconfig:"MYSQL_PORT"`
	MysqlDatabase              string   `envconfig:"MYSQL_DATABASE"`
	MysqlLogLevel              int      `envconfig:"MYSQL_LOG_LEVEL"`
	MysqlMaxIdleConns          int      `envconfig:"MYSQL_MAX_IDLE_CONNS"`
	MysqlMaxOpenConns          int      `env:"MYSQL_MAX_OPEN_CONNS,required"`
	MysqlConnectionMaxLifetime int      `env:"MYSQL_CONNECTION_MAX_LIFETIME,required"`
	LogMaskKeys                []string `env:"LOG_MASK_KEYS" envSeparator:","`
}

func (e *AppEnvironment) ginMode(m string) string {
	// 値補正
	if m != gin.DebugMode && m != gin.ReleaseMode {
		m = gin.DebugMode
	}
	return m
}

func LoadEnvironment() (*AppEnvironment, error) {
	var env AppEnvironment
	err := envconfig.Process("", &env)
	if err != nil {
		return nil, err
	}

	// 値補正
	env.AppMode = env.ginMode(env.AppMode)

	return &env, nil
}

func LoadEnvironmentFromDotenv() (*AppEnvironment, error) {
	err := LoadDotEnv()
	if err != nil {
		return nil, err
	}
	return LoadEnvironment()
}

func LoadDotEnv() error {
	return godotenv.Load()
}
