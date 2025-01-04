package handlers

import (
	"github.com/gin-gonic/gin"

	"ginapp/internal/app"
)

// RegisterRouteHandler is a function to register route handler
func RegisterRouteHandler(app *app.App) {
	// curl -i http://localhost:8080/ping
	app.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := app.Engine.Group("/api")
	// curl -i http://localhost:8080/api/hello
	api.GET("/hello", Hello())

	// curl -X POST -H "Content-Type: application/json" -d "{"name" : "佐藤" , "mail" : "sato@example.com"}" localhost:8080/api/account
	api.POST("/account", PostAccount(app.Mysql))
}
