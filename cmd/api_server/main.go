package main

import (
	"fmt"
	"ginapp/infra/api_server"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	a := &api_server.Config{}
	fmt.Println(a.Host)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pongs",
		})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
