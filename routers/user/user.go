package user

import (
	"github.com/gin-gonic/gin"
	"go_web/controller/user"
)

func RegisterUserRoutes(r *gin.Engine) {
	RegisterUserInfoRouter(r)
	RegisterOrgListRouter(r)
	RegisterTableDataRouter(r)
}

func RegisterUserInfoRouter(r *gin.Engine) {
	r.GET("/userInfo", user.GetUserInfoHandler)
}

func RegisterOrgListRouter(r *gin.Engine) {
	r.GET("/orgList", user.GetUserOrgListHandle)
}

func RegisterTableDataRouter(r *gin.Engine) {
	r.GET("/userDataList", user.GetTableDataForOrgHandler)
}
