package route

import (
	"team_todo/api/v1/controller"
	//"team_todo/api/v1/middleware"

	"github.com/gin-gonic/gin"
)

// 设置路由
// 登录、注册、找回密码...见api文档
func SetupRoutes(r *gin.Engine) {

	userController := &controller.UserController{}
	//登录
	r.GET("/users/login", userController.ShowLoginPage)
	r.POST("/users/login", userController.Login)
	//获取用户资料
	r.GET("/userss/:id/profile", userController.GetProfile)
	//上传头像
	r.POST("/avatars", userController.UploadAvatar)
	//注册
	r.POST("/users/register", controller.Register)

	//更新
	r.PUT("/user/profile", controller.Update)

	//发送邮箱验证码
	r.GET("/user/verify-code", controller.SendVerCodeByEmail)

	//运用中间件的路由
	// User_Auth_Group := r.Group("/users")
	// User_Auth_Group.Use(middleware.AuthMiddleware())
	// {
	// 	User_Auth_Group.GET("/:id/profile", userController.GetProfile)
	// 	User_Auth_Group.PUT("/profile", controller.Update)
	// 	User_Auth_Group.GET("/verify-code", controller.SendVerCodeByEmail)

	// }
	//重设密码(注册时发送“发送邮箱验证码”“重设密码”“注册”请求)

	r.GET("/users/verify-code", controller.ResetPassword)
	groupController := &controller.GroupController{}
	//创建群组
	r.POST("/groups", groupController.CreateGroup)

	//获取已加入群组列表
	r.GET("/groups", groupController.GetGroupList)

	//获取群组信息
	r.GET("/groups/:id/info", groupController.GetGroupInfo)

	//获取群组成员
	r.GET("/groups/:id/members", groupController.GetGroupMembers)

	//加入群组
	r.GET("/groups/join", groupController.JoinGroup)

	//获取群组邀请码
	r.GET("groups/:id/code", groupController.GetGroupCode)

	//退出群组
	r.POST("/groups/:id/leave", groupController.LeaveGroup)

	//解散群组
	r.DELETE("/groups/:id", groupController.DelGroup)

}
