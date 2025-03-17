package routers

import (
	"github.com/gin-gonic/gin"
	"go_web/middleware"
	"go_web/routers/login"
	"go_web/routers/register"
	"go_web/routers/user"
)

// InitRouter 初始化所有路由
func InitRouter() *gin.Engine {
	r := gin.Default()
	// 跨域  middleware.AuthMiddleware()
	r.Use(gin.Logger(), gin.Recovery(), middleware.CORS())
	// 注册 login 相关的路由
	login.RegisterLoginRouter(r)
	register.RegisterUserRouter(r)
	user.RegisterUserInfoRouter(r)
	return r
}
