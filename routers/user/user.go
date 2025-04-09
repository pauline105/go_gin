package user

import (
	"github.com/gin-gonic/gin"
	"go_web/controller/user"
)

func RegisterUserRoutes(r *gin.Engine) {
	RegisterUserInfoRouter(r)
	RegisterOrgListRouter(r)
}

func RegisterUserInfoRouter(r *gin.Engine) {
	r.GET("/userInfo", user.GetUserInfoHandler)
}

func RegisterOrgListRouter(r *gin.Engine) {
	r.GET("/orgList", user.GetUserInfoHandle)
}
