package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"ginapp/internal/app"
)

func main() {
	env, _ := app.LoadEnvironmentFromDotenv()
	r := gin.Default()
	fmt.Println(env.AppName)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongs",
		})
	})
	r.Run(fmt.Sprintf("%s:%d", env.AppHost, env.AppPort)) // 0.0.0.0:8080 でサーバーを立てます。
}
