package app

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	AppName    string `envconfig:"APP_NAME"`
	AppHost    string `envconfig:"APP_HOST"`
	AppPort    int    `envconfig:"APP_PORT"`
	AppVersion string `envconfig:"APP_VERSION"`
	AppMode    string `envconfig:"APP_MODE"`
}

func LoadEnvironment() (*Environment, error) {
	var env Environment
	err := envconfig.Process("", &env)
	if err != nil {
		return nil, err
	}
	return &env, nil
}

func LoadEnvironmentFromDotenv() (*Environment, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return LoadEnvironment()
}
