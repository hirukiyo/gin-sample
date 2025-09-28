package environment

import (
	"errors"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

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

func LoadAppEnvironment() (*AppEnvironment, error) {
	return loadEnvironmentFromDotenv()
}

func (e *AppEnvironment) ginMode(m string) string {
	// 値補正
	if m != gin.DebugMode && m != gin.ReleaseMode {
		m = gin.DebugMode
	}
	return m
}

func loadEnvironment() (*AppEnvironment, error) {
	var env AppEnvironment
	err := envconfig.Process("", &env)
	if err != nil {
		slog.Info("binding env param error.", "err", err)
		return nil, err
	}

	// 値補正
	env.AppMode = env.ginMode(env.AppMode)

	return &env, nil
}

func loadEnvironmentFromDotenv() (*AppEnvironment, error) {
	path, _ := getEnvFilePath()
	err := godotenv.Load(path)
	if err != nil {
		slog.Info("dotenv load error.", "err", err)
		return nil, err
	}
	return loadEnvironment()
}

func getEnvFilePath() (string, error) {
	// プログラムのカレントディレクトリーを探索起点にする
	dir, _ := os.Getwd()

	for {
		// 探索ディレクトリーに.envファイルが存在するか確認する
		envPath := filepath.Join(dir, ".env")
		if info, err := os.Stat(envPath); err == nil && !info.IsDir() {
			return envPath, nil
		}

		// 探索ディレクトリーを親ディレクトリーへ移動させる
		parentDir := filepath.Dir(dir)
		if parentDir == dir {
			return "", errors.New(".env file was not found")
		}

		dir = parentDir
	}
}
