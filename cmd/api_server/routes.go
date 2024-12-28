package main

import (
	"ginapp/internal/handlers"

	"github.com/gin-gonic/gin"
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
