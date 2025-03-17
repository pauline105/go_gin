package user

import (
	"github.com/gin-gonic/gin"
	"go_web/controller/user"
)

func RegisterUserInfoRouter(r *gin.Engine) {
	r.GET("/userInfo", user.GetUserInfoHandler)
}
