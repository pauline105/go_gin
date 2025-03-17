package jwt_plugin

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

var key = "waley326"

// Data 创建token JWT 內部的 Payload
type Data struct {
	Username string
	Password string
	Id       int
	jwt.RegisteredClaims
}

// Sign 生成token
func Sign(data jwt.Claims) (string, error) {
	//	jwt.NewWithClaims(...)：創建一個新的 JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	//	使用 key（密鑰）對 JWT 進行簽名
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return "", nil
	}

	return tokenString, nil
}

// ValidateToken 解析JWT
func ValidateToken(tokenString string) (*Data, error) {
	if strings.HasPrefix(tokenString, "Bearer ") {
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	}
	// 如果是自定义Claim结构体则需要使用 ParseWithClaims 方法
	token, err := jwt.ParseWithClaims(tokenString, &Data{}, func(token *jwt.Token) (i interface{}, err error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// 直接使用标准的Claim则可以直接使用Parse方法
		//token, err := jwt.Parse(tokenString, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(key), nil
	})

	if claims, ok := token.Claims.(*Data); ok && token.Valid {
		//a := token.Valid
		//b := claims.Valid()
		//fmt.Println(a, b)
		return claims, nil
	} else {
		return nil, err
	}
}
