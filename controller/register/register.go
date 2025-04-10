package register

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_web/db"
	"log"

	"go_web/type/structType/login"
)

// RegisterUserHandler 註冊用戶邏輯
func RegisterUserHandler(c *gin.Context) {
	//	獲取資料庫實例
	sqlxDB := db.ConnectDB()
	// 	綁定用戶數據類型
	RegisterUserData := new(login.LoginType)
	//	獲取用戶註冊數據
	if err := c.ShouldBindJSON(&RegisterUserData); err != nil {
		c.JSON(400, gin.H{
			"code": 400,
			"info": "请求格式错误或缺少必填字段",
		})
		return
	}
	fmt.Println(RegisterUserData)
	// 資料庫查找是否有相同的用戶名
	var count int
	sqlString := "SELECT COUNT(*) FROM user WHERE username = ?"
	err := sqlxDB.Get(&count, sqlString, RegisterUserData.Username)
	if err != nil {
		log.Fatalln(err)
	}
	// 找到相同用戶名
	if count > 0 {
		c.JSON(409, gin.H{
			"code":    409,
			"message": "用戶名已存在",
		})
		return
	}
	// 先給資料表添加數據獲取到id
	addUserProfile := "INSERT INTO profile (name,status) VALUES (?,?)"
	res, err := sqlxDB.Exec(addUserProfile, "", 1)
	if err != nil {
		log.Fatalln(err)
		return
	}
	// 獲取用戶插入的profile_id
	profileId, err := res.LastInsertId()
	// 給密碼hash加密
	hashedPassword, err := db.HashPassword(RegisterUserData.Password)
	// 用戶存入user表
	query := "INSERT INTO user (username,password,profile_id) VALUES (?,?,?)"
	_, err = sqlxDB.Exec(query, RegisterUserData.Username, hashedPassword, profileId)
	if err != nil {
		log.Fatalf("Error inserting user: %v", err)
		return
	}
	fmt.Printf("Parsed Data: %+v\n", RegisterUserData)
	c.JSON(200, gin.H{
		"status":  200,
		"message": "用戶註冊成功",
	})
}
