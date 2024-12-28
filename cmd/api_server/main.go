package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/MatusOllah/slogcolor"
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

	setLogger(env)
	gin.SetMode(env.AppMode)

	app := &App{
		cfg: env,
		gin: gin.Default(),
	}

	slog.Info("app start.", "env", app.cfg)
	applog.Info(context.Background(), "app end.", "env", app.cfg)

	app.gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongs",
		})
	})

	app.gin.Run(fmt.Sprintf("%s:%d", app.cfg.AppHost, app.cfg.AppPort))

}

func setLogger(cfg *app.AppEnvironment) {
	var h slog.Handler
	logLevel := logLevel(cfg.AppLogLevel)
	if cfg.AppMode == gin.ReleaseMode {
		h = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true, Level: logLevel})
	} else {
		h = slogcolor.NewHandler(os.Stdout, &slogcolor.Options{
			Level:       logLevel,
			TimeFormat:  time.RFC3339,
			SrcFileMode: slogcolor.ShortFile,
		})
	}
	slog.SetDefault(slog.New(h))
}

func logLevel(l string) slog.Level {
	appLogLevel := slog.LevelInfo
	switch strings.ToLower(l) {
	case "debug":
		appLogLevel = slog.LevelDebug
	case "info":
		appLogLevel = slog.LevelInfo
	case "warn":
		appLogLevel = slog.LevelWarn
	case "error":
		appLogLevel = slog.LevelError
	}
	return appLogLevel
}
