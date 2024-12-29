package main

import (
	"github.com/gin-gonic/gin"

	"ginapp/internal/handlers"
)

// RegisterRouteHandler is a function to register route handler
func RegisterRouteHandler(app *App) {
	app.gin.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := app.gin.Group("/api")
	api.GET("/hello", handlers.Hello())
}
