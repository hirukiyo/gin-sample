package handlers

import (
	"github.com/gin-gonic/gin"

	"ginapp/internal/app/applog"
)

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		applog.Debug(c, "execute hello handler")
		c.JSON(200, gin.H{
			"message": "hello",
		})
	}
}
