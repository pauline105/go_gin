package login

import (
	"github.com/gin-gonic/gin"
	"go_web/controller/login"
)

// RegisterLoginRouter 注册 home 相关的路由
func RegisterLoginRouter(r *gin.Engine) {
	r.POST("/login", login.LoginHandle)
}
