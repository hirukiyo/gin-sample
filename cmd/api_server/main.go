package main

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"

	"ginapp/database/mysql"
	"ginapp/internal/app"
	"ginapp/internal/app/applog"
)

func main() {
	env, err := app.LoadEnvironmentFromDotenv()
	if err != nil {
		slog.Error("environment load error.", "err", err)
		os.Exit(1)
	}

	applog.SetLogger(env)
	gin.SetMode(env.AppMode)
	engine := gin.Default()

	// middleware
	{
		engine.Use(gin.Recovery())
		engine.Use(RequestLoggerMiddleware())
	}

	// setup
	mysqlConn, err := mysql.NewConnection(
		env.MysqlUser,
		env.MysqlPassword,
		env.MysqlHost,
		env.MysqlPort,
		env.MysqlDatabase,
		env.MysqlLogLevel,
		env.MysqlMaxIdleConns,
		env.MysqlMaxOpenConns,
		env.MysqlConnectionMaxLifetime,
	)
	if err != nil {
		slog.Error("mysql connection error.", "err", err)
		os.Exit(1)
	}

	// create app
	app := &app.App{
		Env:    env,
		Engine: engine,
		Mysql:  mysqlConn,
	}

	slog.Info("app start.", "env", app.Env)

	RegisterRouteHandler(app)

	if err := app.Engine.Run(fmt.Sprintf("%s:%d", app.Env.AppHost, app.Env.AppPort)); err != nil {
		slog.Error("server run error.", "err", err)
		os.Exit(1)
	}
}

func RequestLoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := xid.New().String()

		// extracting request body
		body := []byte{}
		if c.Request.Method == http.MethodPost || c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPatch {
			body, _ = io.ReadAll(c.Request.Body)
			c.Request.Body = io.NopCloser(bytes.NewReader(body))
		}

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
			"body", string(body),
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
