package db

import (
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB 全局注冊數據庫實例
var DB *sqlx.DB

func ConnectDB() *sqlx.DB {
	// 數據庫地址
	dsn := "root:123456@tcp(127.0.0.1:3306)/ww?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	// 創建gorm數據庫實例
	DB, err = sqlx.Connect("mysql", dsn)

	if err != nil {
		log.Fatalf("鏈接數據庫失敗: %v", err)
	}

	// 設置每個連接的最大生命周期
	DB.SetConnMaxLifetime(time.Second * 10)

	// 數據庫最大連接數
	DB.SetMaxOpenConns(200)

	// 數據庫最大空閑連接數
	DB.SetMaxIdleConns(10)

	return DB
}

// GetPasswordHashByUsername 根据用户名查询用户的哈希密码
func GetPasswordHashByUsername(db *sqlx.DB, username string) (string, error) {
	var passwordHash string
	query := "SELECT password FROM user WHERE username = ?"
	err := db.QueryRow(query, username).Scan(&passwordHash)
	return passwordHash, err
}

// HashPassword 注冊用戶的時候 給用戶密碼設置hash加密
func HashPassword(password string) (string, error) {
	// Generate a hash for the password, with a cost factor of 10
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
