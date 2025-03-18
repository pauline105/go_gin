package login

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go_web/db"
	"go_web/jwt_plugin"
	"go_web/type/structType/login"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

func LoginHandle(c *gin.Context) {
	// 獲取數據庫實例
	sqlxDB := db.ConnectDB()
	loginData := new(login.LoginType)

	// 獲取接收到的數據
	if err := c.ShouldBindJSON(loginData); err != nil {

		c.JSON(400, gin.H{
			"error": "Invalid request data",
		})
		return
	}

	// 獲取到用戶的哈希密碼
	storedHash, storedHashErr := db.GetPasswordHashByUsername(sqlxDB, loginData.Username)
	// 如果找不到用戶 給前端返回找不到用戶的信息
	if storedHashErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "401",
			"message": "找不到用戶",
		})
		return
	}
	// 判斷用戶輸入密碼是否和數據庫一致
	storedStatus := bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(loginData.Password))
	if storedStatus != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "401",
			"message": "hash密碼不一樣",
		})
		return
	}

	// 獲取到當前登錄用戶信息
	var user login.User
	sqlStr := "SELECT id, username, password, profile_id FROM user WHERE username = ? AND password = ?"
	err := sqlxDB.Get(&user, sqlStr, loginData.Username, storedHash)

	if err != nil {
		return
	}

	data := jwt_plugin.Data{
		Username: loginData.Username,
		Password: loginData.Password,
		Id:       user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			//  过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 12)),
			//  签发时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			//	生效时间
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	sign, err := jwt_plugin.Sign(data)

	c.JSON(http.StatusOK, gin.H{
		"token":   "Bearer " + sign,
		"message": "登錄成功",
		"status":  200,
	})
}
