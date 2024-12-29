package handlers

import (
	"github.com/gin-gonic/gin"

	"ginapp/database/models"
	"ginapp/database/mysql"
	"ginapp/internal/app/applog"
)

// curl -X POST -H "Content-Type: application/json" -d "{"name" : "佐藤" , "mail" : "sato@example.com"}" localhost:8080/api/account
func PostAccount(db *mysql.Connection) gin.HandlerFunc {
	return func(c *gin.Context) {
		applog.Debug(c, "execute PostAccount handler")
		// account := &models.Account{
		// 	Name:     "Test1",
		// 	Email:    "test@email.com",
		// 	Password: "ffff",
		// }
		// if result := db.Create(&account); result.Error != nil {
		// 	applog.Error(c, "account create error", "err", result.Error)
		// 	c.JSON(500, gin.H{
		// 		"message": "Internal Server Error",
		// 	})
		// 	return
		// }

		var account models.Account
		if err := db.First(&account).Error; err != nil {
			applog.Error(c, "account fetch error", "err", err)
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
		applog.Info(c, "account create success", "account", account)
		c.JSON(200, gin.H{
			"message": "PostAccount",
		})
	}
}
