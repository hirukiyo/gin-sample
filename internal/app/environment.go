package app

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppEnvironment struct {
	AppMode     string `envconfig:"APP_MODE"`
	AppName     string `envconfig:"APP_NAME"`
	AppHost     string `envconfig:"APP_HOST"`
	AppPort     int    `envconfig:"APP_PORT"`
	AppVersion  string `envconfig:"APP_VERSION"`
	AppLogLevel string `envconfig:"APP_LOG_LEVEL"`
	AppTest     string `envconfig:"APP_TEST"`
}

func LoadEnvironment() (*AppEnvironment, error) {
	var env AppEnvironment
	err := envconfig.Process("", &env)
	if err != nil {
		return nil, err
	}

	// 値補正
	env.AppMode = ginMode(env.AppMode)

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

func ginMode(m string) string {
	// 値補正
	if m != gin.DebugMode && m != gin.ReleaseMode {
		m = gin.DebugMode
	}
	return m
}
