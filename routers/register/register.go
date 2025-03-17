package register

import (
	"github.com/gin-gonic/gin"
	"go_web/controller/register"
)

// RegisterUserRouter 注册相关的路由
func RegisterUserRouter(r *gin.Engine) {
	r.POST("/register", register.RegisterUserHandler)
}
