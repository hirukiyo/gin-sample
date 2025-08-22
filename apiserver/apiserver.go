package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/samber/lo"

	"github.com/hirukiyo/gin-sample/apiserver/app"
	"github.com/hirukiyo/gin-sample/apiserver/applog"
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

// Middleware to log requests
func LoggingMiddleware(maskKeys []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody interface{}
		bodyBytes, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // Restore the body for the handler

		contentType := c.GetHeader("Content-Type")
		if strings.Contains(contentType, "application/json") {
			if err := json.Unmarshal(bodyBytes, &requestBody); err == nil {
				requestBody = maskAndTruncateJSON(requestBody, maskKeys)
			} else {
				requestBody = string(bodyBytes)
			}
		} else {
			requestBody = string(bodyBytes)
		}

		log := fmt.Sprintf("Method: %s, Path: %s, Query: %s, Body: %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.Request.URL.RawQuery,
			requestBody,
		)
		fmt.Println(log)

		c.Next()
	}
}

func maskAndTruncateJSON(data interface{}, maskKeys []string) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{})
		for key, value := range v {
			if lo.Contains(maskKeys, key) {
				result[key] = "***MASKED***"
			} else {
				result[key] = maskAndTruncateJSON(value, maskKeys)
			}
		}
		return result
	case []interface{}:
		for i, value := range v {
			v[i] = maskAndTruncateJSON(value, maskKeys)
		}
		return v
	case string:
		if len(v) > 256 {
			return v[:256]
		}
		return v
	default:
		return v
	}
}
