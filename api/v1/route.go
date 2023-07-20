package route

import (
	"team_todo/api/v1/controller"

	"github.com/gin-gonic/gin"
)

// 设置路由
// 登录、注册、找回密码...见api文档
func SetupRoutes(r *gin.Engine) {

	userController := &controller.UserController{}
	//登录
	r.GET("/user/login", userController.ShowLoginPage)
	r.POST("/user/login", userController.Login)
	r.GET("GET /users/:id/profile", userController.GetProfile)

	//注册
	r.POST("/user/register", controller.Register)

	//更新
	r.PUT("/user/profile", controller.Update)

	//发送邮箱验证码
	r.GET("/user/verify-code", controller.SendVerCodeByEmail)

	//运用中间件的路由
	// User_Auth_Group := r.Group("/user" , middleware.AuthMiddleware())
	// {
	// 	User_Auth_Group.GET("/:id/profile", userController.GetProfile)
	// 	User_Auth_Group.PUT("/profile",  controller.Update)

	// }
	//重设密码(注册时发送“发送邮箱验证码”“重设密码”“注册”请求)

	r.POST("/users/resetpassword", controller.ResetPassword)
	groupController := &controller.GroupController{}
	//创建群组
	r.POST("/groups", groupController.CreateGroup)

}
