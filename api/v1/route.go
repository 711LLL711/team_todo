package route

import (
	"team_todo/controller"

	"github.com/gin-gonic/gin"
)

// 设置路由
// 登录、注册、找回密码...见api文档
func SetupRoutes(r *gin.Engine) {
	userController := controller.UserController{}
	r.GET("/user/login", userController.ShowLoginPage)
	r.POST("/user/login", userController.Login)
}
