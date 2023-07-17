package route

import (
	"team_todo/api/v1/controller"

	"github.com/gin-gonic/gin"
)

// 设置路由
// 登录、注册、找回密码...见api文档
func SetupRoutes(r *gin.Engine) {
	userController := &controller.UserController{}
	r.GET("/user/login", userController.ShowLoginPage)
	r.POST("/user/login", userController.Login)
	r.GET("GET /users/:id/profile", userController.GetProfile)

	groupController := &controller.GroupController{}
	r.POST("/groups", groupController.CreateGroup)
	r.GET("/groups", groupController.GetGroupList)
	r.GET("groups/:id/info", groupController.GetGroupInfo)
	r.GET("groups/:id/members", groupController.GetGroupMembers)
}
