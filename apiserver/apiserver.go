package apiserver

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"

	"github.com/hirukiyo/gin-sample/apiserver/app"
	"github.com/hirukiyo/gin-sample/apiserver/applog"
	"github.com/hirukiyo/gin-sample/apiserver/middleware"
	"github.com/hirukiyo/gin-sample/infra/mysql"
	"github.com/hirukiyo/gin-sample/internal/handlers"
)

func StartAPIServer() int {
	env, err := app.LoadEnvironmentFromDotenv()
	if err != nil {
		slog.Error("environment load error.", "err", err)
		return 1
	}

	applog.SetLogger(&applog.SetLoggerInput{
		AppMode:     env.AppMode,
		AppLogLevel: env.AppLogLevel,
	})
	gin.SetMode(env.AppMode)
	engine := gin.Default()

	// middleware
	{
		engine.Use(gin.Recovery())
		engine.Use(middleware.RequestLoggingMiddleware([]string{"password", "secret"}))
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
		return 1
	}

	// create app
	app := &app.App{
		Env:    env,
		Engine: engine,
		Mysql:  mysqlConn,
	}

	slog.Info("app start.", "env", app.Env)

	handlers.RegisterRouteHandler(app)

	if err := app.Engine.Run(fmt.Sprintf("%s:%d", app.Env.AppHost, app.Env.AppPort)); err != nil {
		slog.Error("server run error.", "err", err)
		return 1
	}

	return 0
}
