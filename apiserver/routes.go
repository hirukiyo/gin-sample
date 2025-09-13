package apiserver

import (
	"github.com/gin-gonic/gin"

	"github.com/hirukiyo/gin-sample/apiserver/handlers"
	"github.com/hirukiyo/gin-sample/application/usecase"
)

// RegisterRouteHandler is a function to register route handler
func RegisterRouteHandler(app *App) {
	// curl -i http://localhost:8080/ping
	app.Engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	api := app.Engine.Group("/api")
	// curl -i http://localhost:8080/api/hello
	api.GET("/hello", handlers.Hello())

	accountUsecase := usecase.NewAccountUsecase(app.GormDB)

	// curl -X POST -H "Content-Type: application/json" -d '{"name":"test user 1", "email":"test_user_1@example.com", "password":"password"}' localhost:8080/api/accounts
	api.POST("/accounts", handlers.PostAccount(app.GormDB, accountUsecase))

	// curl -X GET -H "Content-Type: application/json" localhost:8080/api/accounts
	api.GET("/accounts", handlers.FindAccounts(app.GormDB, accountUsecase))

	// curl -X GET -H "Content-Type: application/json" localhost:8080/api/accounts/:id
	api.GET("/accounts/:id", handlers.FindAccountByID(app.GormDB, accountUsecase))

	// curl -X PUT -H "Content-Type: application/json" localhost:8080/api/accounts/:id
	api.PUT("/accounts/:id", handlers.UpdateAccountByID(app.GormDB, accountUsecase))

	// curl -X DELETE -H "Content-Type: application/json" localhost:8080/api/accounts/:id
	api.DELETE("/accounts/:id", handlers.DeleteAccountByID(app.GormDB, accountUsecase))
}
