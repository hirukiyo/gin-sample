package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	"ginapp/internal/app"
	"ginapp/internal/app/applog"
)

type App struct {
	cfg *app.AppEnvironment
	gin *gin.Engine
}

func main() {
	env, err := app.LoadEnvironmentFromDotenv()
	if err != nil {
		slog.Error("environment load error.", "err", err)
		os.Exit(1)
	}

	applog.SetLogger(env)
	gin.SetMode(env.AppMode)

	app := &App{
		cfg: env,
		gin: gin.Default(),
	}

	slog.Info("app start.", "env", app.cfg)
	applog.Info(context.Background(), "app end.", "env", app.cfg)

	RegistRouteHandler(app)

	app.gin.Run(fmt.Sprintf("%s:%d", app.cfg.AppHost, app.cfg.AppPort))

}
