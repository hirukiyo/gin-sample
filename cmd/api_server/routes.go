package main

import (
	"github.com/gin-gonic/gin"

	"ginapp/internal/app"
	"ginapp/internal/handlers"
)

// RegisterRouteHandler is a function to register route handler
func RegisterRouteHandler(app *app.App) {
	app.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := app.Engine.Group("/api")
	api.GET("/hello", handlers.Hello())

	api.POST("/account", handlers.PostAccount(app.Mysql))
}
