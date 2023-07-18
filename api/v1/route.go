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

	//注册
	r.POST("/user/register", controller.Register)

	//更新
	r.PUT("/user/profile") //鉴权中间件,controller.Update

	//发送邮箱验证码
	r.GET("/user/verify-code", controller.SendVerCodeByEmail)

	//重设密码(注册时发送“发送邮箱验证码”“重设密码”“注册”请求)

	r.GET("/users/verify-code", controller.ResetPassword)
	groupController := &controller.GroupController{}
	r.POST("/groups", groupController.CreateGroup)
	r.GET("/groups", groupController.GetGroupList)
	r.GET("groups/:id/info", groupController.GetGroupInfo)
	r.GET("groups/:id/members", groupController.GetGroupMembers)

}
