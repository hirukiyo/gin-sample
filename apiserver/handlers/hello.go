package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hirukiyo/gin-sample/apiserver/applog"
)

func Hello() gin.HandlerFunc {
	return func(c *gin.Context) {
		applog.Debug(c, "execute hello handler")
		c.JSON(http.StatusOK, gin.H{
			"message": "hello",
		})
	}
}
