package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"

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

	// middleware
	{
		app.gin.Use(gin.Recovery())
		app.gin.Use(RequestLoggerMiddleware())
	}

	slog.Info("app start.", "env", app.cfg)
	applog.Info(context.Background(), "app end.", "env", app.cfg)

	RegisterRouteHandler(app)

	if err := app.gin.Run(fmt.Sprintf("%s:%d", app.cfg.AppHost, app.cfg.AppPort)); err != nil {
		slog.Error("server run error.", "err", err)
		os.Exit(1)
	}
}

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := xid.New().String()

		applog.Info(c, "REQ",
			"id", id,
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"query", c.Request.URL.RawQuery,
			"client_ip", c.ClientIP(),
			"real_ip", c.GetHeader("X-Real-IP"),
			"forward-for", c.GetHeader("X-Forwarded-For"),
			"host", c.Request.Host,
			"user_agent", c.GetHeader("User-Agent"),
			"referer", c.GetHeader("Referer"),
		)

		startTime := time.Now()
		c.Next()
		latency := time.Since(startTime)

		applog.Info(c, "RES",
			"id", id,
			"status", c.Writer.Status(),
			"size", c.Writer.Size(),
			"latency", slog.AnyValue(latency),
		)
	}
}
