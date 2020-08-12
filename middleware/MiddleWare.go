package middleware

import (
	"ginEssential/common"
	"ginEssential/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func AuthMiddleWare() gin.HandlerFunc{
	return func(c *gin.Context) {
		// 获取授权的header
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 400,
				"msg": "权限不足",
			})
			c.Abort()
			return

		}
		// 截取token
		tokenString = tokenString[7:]

		token, claim, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 400,
				"msg": "权限不足",
			})
			c.Abort()
			return
		}

		// 获取claim中的userId
		userId := claim.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)
		// 验证用户id
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 400,
				"msg": "权限不足",
			})
			c.Abort()
			return
		}

		// 用户存在，写入上下文
		c.Set("user", user)
		c.Next()

	}

}