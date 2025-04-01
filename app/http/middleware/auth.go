package middleware

import (
	"context"
	"net/http"

	"347613781qq.com/goInit1/app/global"
	"347613781qq.com/goInit1/app/service"
	"347613781qq.com/goInit1/utils"
	"github.com/gin-gonic/gin"
)

type HeaderParams struct {
	Authorization string `header:"Authorization" binding:"required,min=20"`
}

func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "未提供 Token", "code": global.LoginErrorCode})
			c.Abort()
			return
		}

		// 检查是否在黑名单
		ctx := context.Background()
		if utils.Redis.Exists(ctx, "blacklist:"+token).Val() == 1 {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token 已失效，请重新登录", "code": global.LoginErrorCode})
			c.Abort()
			return
		}

		// 解析 Token
		claims, err := service.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "Token 无效", "code": global.LoginErrorCode})
			c.Abort()
			return
		}

		// 传递用户名到上下文
		c.Set("username", claims.Username)
		c.Next()

	}

}
