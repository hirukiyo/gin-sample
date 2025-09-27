package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/hirukiyo/gin-sample/apiserver/applog"
	"github.com/hirukiyo/gin-sample/application/usecases"
	"github.com/hirukiyo/gin-sample/domain"
	"github.com/hirukiyo/gin-sample/infra/mysql/models"
)

// curl -X POST -H "Content-Type: application/json" -d '{"name":"test user 1", "email":"test_user_1@example.com", "password":"password"}' localhost:8080/api/accounts
func PostAccount(db *gorm.DB, uc usecases.AccountUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		applog.Debug(c, "execute PostAccount handler")
		var account models.Account
		if err := c.ShouldBindJSON(&account); err != nil {
			applog.Error(c, "invalid request body", "err", err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		err := gorm.G[models.Account](db).Create(c, &account)
		if err != nil {
			applog.Error(c, "account create error", "err", err)
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}
		c.JSON(201, gin.H{
			"result": account,
		})
	}
}

// curl -X GET -H "Content-Type: application/json" localhost:8080/api/accounts
func FindAccounts(db *gorm.DB, uc usecases.AccountUsecase) gin.HandlerFunc {
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

// curl -i -X GET -H "Content-Type: application/json" localhost:8080/api/accounts/1
func GetAccount(db *gorm.DB, uc usecases.AccountUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		applog.Debug(c, "execute GetAccount handler")

		paramId := c.Param("id")
		// idが未指定の場合は400を返却
		if paramId == "" {
			applog.Warn(c, "id is not specified")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "id is not specified",
			})
			return
		}

		id, err := strconv.ParseUint(paramId, 10, 64)
		if err != nil {
			applog.Warn(c, "id is invalid parameter", "id", paramId)
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "id is invalid parameter",
			})
			return
		}

		account, err := uc.GetAccount(c, id)
		if err != nil {
			// idが存在しない場合は404を返却
			if errors.Is(err, domain.ErrNotFound) {
				applog.Warn(c, "account not found", "id", id, "err", err)
				c.JSON(http.StatusNotFound, gin.H{
					"message": "Not Found",
				})
				return
			}
			// その他のエラーは500を返却
			applog.Error(c, "account fetch error", "err", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		// 取得できた場合は200とaccount情報を返却
		applog.Debug(c, "account fetch success", "account", account)
		c.JSON(http.StatusOK, gin.H{
			"result": account,
		})
	}
}

// curl -X DELETE -H "Content-Type: application/json" localhost:8080/api/accounts/1
func DeleteAccount(db *gorm.DB, uc usecases.AccountUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		applog.Debug(c, "execute DeleteAccount handler")

		id := c.Param("id")
		// idが未指定の場合は400を返却
		if id == "" {
			applog.Warn(c, "id is not specified")
			c.JSON(400, gin.H{
				"message": "id is not specified",
			})
			return
		} else if _, err := strconv.ParseUint(id, 10, 64); err != nil {
			applog.Warn(c, "id is not uint64", "id", id)
			c.JSON(400, gin.H{
				"message": "id is not uint64",
			})
			return
		}

		rowsAffected, err := gorm.G[models.Account](db).Where("id = ?", id).Delete(c)
		if err != nil {
			// その他のエラーは500を返却
			applog.Error(c, "account delete error", "err", err)
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		if rowsAffected == 0 {
			// idが存在しない場合は404を返却
			applog.Warn(c, "account not found", "id", id)
			c.JSON(404, gin.H{
				"message": "Not Found",
			})
			return
		}

		// 削除できた場合は200を返却
		applog.Debug(c, "account delete success", "id", id)
		c.JSON(200, nil)
	}
}

// curl -X PUT -H "Content-Type: application/json" -d '{"name":"new_name", "email":"new@example.com", "password":"new_password"}' localhost:8080/api/accounts/1
func UpdateAccount(db *gorm.DB, uc usecases.AccountUsecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		applog.Debug(c, "execute UpdateAccount handler")

		id := c.Param("id")
		// idが未指定の場合は400を返却
		if id == "" {
			applog.Warn(c, "id is not specified")
			c.JSON(400, gin.H{
				"message": "id is not specified",
			})
			return
		} else if _, err := strconv.ParseUint(id, 10, 64); err != nil {
			applog.Warn(c, "id is not uint64", "id", id)
			c.JSON(400, gin.H{
				"message": "id is not uint64",
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

		var req models.Account
		if err := c.ShouldBindJSON(&req); err != nil {
			applog.Error(c, "invalid request body", "err", err)
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		account.Name = req.Name
		account.Email = req.Email
		account.Password = req.Password

		rowsAffected, err := gorm.G[models.Account](db).Updates(c, account)
		if err != nil {
			// その他のエラーは500を返却
			applog.Error(c, "account update error", "err", err)
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		if rowsAffected == 0 {
			// ここでrowAffectedが0の場合は更新が失敗している
			applog.Error(c, "account update row affected is zero", "err", err)
			c.JSON(500, gin.H{
				"message": "Internal Server Error",
			})
			return
		}

		c.JSON(200, gin.H{
			"result": account,
		})
	}
}
