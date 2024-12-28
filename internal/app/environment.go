package app

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type AppEnvironment struct {
	AppName    string `envconfig:"APP_NAME"`
	AppHost    string `envconfig:"APP_HOST"`
	AppPort    int    `envconfig:"APP_PORT"`
	AppVersion string `envconfig:"APP_VERSION"`
	AppMode    string `envconfig:"APP_MODE"`
}

func LoadEnvironment() (*AppEnvironment, error) {
	var env AppEnvironment
	err := envconfig.Process("", &env)
	if err != nil {
		return nil, err
	}
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
