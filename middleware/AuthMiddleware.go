package middleware

import (
	"github.com/gin-gonic/gin"
	"go_web/jwt_plugin"
	"net/http"
)

var whiteList = map[string]bool{
	"/login": true,
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 跳过白名单内的接口
		if _, exists := whiteList[c.Request.URL.Path]; exists {
			c.Next()
			return
		}

		// 检查 Token
		token := c.GetHeader("Authorization")
		_, err := jwt_plugin.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "无效的Token",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
