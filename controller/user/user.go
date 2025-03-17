package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"

	//"go_web/db"
	"go_web/jwt_plugin"
)

// GetUserInfoHandler 獲取用戶信息
func GetUserInfoHandler(c *gin.Context) {
	//sqlxDB := db.ConnectDB()
	token := c.GetHeader("Authorization")
	// 檢驗token 獲取用戶登錄用戶id
	claims, err := jwt_plugin.ValidateToken(token)
	if err != nil {
		log.Fatalf("獲取登錄用戶信息失敗: %v", err)
		return
	}
	// 返回用戶的信息
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
