package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	"ginapp/internal/app"
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

	gin.SetMode(env.AppMode)
	app := &App{
		cfg: env,
		gin: gin.Default(),
	}

	app.gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongs",
		})
	})
	app.gin.Run(fmt.Sprintf("%s:%d", env.AppHost, env.AppPort)) // 0.0.0.0:8080 でサーバーを立てます。
}

func setLogger(app *App) {
	// loggerの設定

}
