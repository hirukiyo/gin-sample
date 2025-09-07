package handlers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/hirukiyo/gin-sample/apiserver/applog"
	"github.com/hirukiyo/gin-sample/infra/mysql/models"
)

// curl -X POST -H "Content-Type: application/json" -d "{"name" : "佐藤" , "mail" : "sato@example.com"}" localhost:8080/api/account
func PostAccount(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		applog.Debug(c, "execute PostAccount handler")
		account := &models.Account{
			Name:     "Test1",
			Email:    "test@email.com",
			Password: "ffff",
		}
		if result := db.Create(&account); result.Error != nil {
			applog.Error(c, "account create error", "err", result.Error)
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		// var account models.Account
		// if err := db.First(&account).Error; err != nil {
		// 	applog.Error(c, "account fetch error", "err", err)
		// 	c.JSON(500, gin.H{
		// 		"message": "Internal Server Error",
		// 	})
		// 	return
		// }
		// applog.Info(c, "account create success", "account", account)
		// c.JSON(200, gin.H{
		// 	"message": account,
		// })
	}
}

// curl -X GET -H "Content-Type: application/json" localhost:8080/api/account
func FindAccounts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		applog.Debug(c, "execute FindAccounts handler")

		accounts, err := gorm.G[models.Account](db).Find(c)
		if err != nil {

			applog.Error(c, "account fetch error", "err", err)
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
		applog.Debug(c, "account fetch success", "accounts", accounts)
		c.JSON(200, gin.H{
			"result": accounts,
		})
	}
}

// curl -X GET -H "Content-Type: application/json" localhost:8080/api/accounts/1
func FindAccountByID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		applog.Debug(c, "execute FindAccountByID handler")

		id := c.Param("id")
		// idが未指定の場合は400を返却
		if id == "" {
			applog.Warn(c, "id is not specified")
			c.JSON(400, gin.H{
				"message": "id is not specified",
			})
			return
		}

		account, err := gorm.G[models.Account](db).Where("id = ?", id).First(c)
		if err != nil {
			// idが存在しない場合は404を返却
			if errors.Is(err, gorm.ErrRecordNotFound) {
				applog.Warn(c, "account not found", "id", id)
				c.JSON(404, gin.H{
					"message": "Not Found",
				})
				return
			}
			// その他のエラーは500を返却
			applog.Error(c, "account fetch error", "err", err)
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		// 取得できた場合は200とaccount情報を返却
		applog.Debug(c, "account fetch success", "account", account)
		c.JSON(200, gin.H{
			"result": account,
		})
	}
}
