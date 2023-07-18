package route

import (
	"team_todo/api/v1/controller"

	"github.com/gin-gonic/gin"
)

//设置路由
//登录、注册、找回密码...见api文档

// 注册、更新、发送邮箱验证码
func MyRouter(r *gin.Engine) {

	//注册
	r.POST("/user/register", controller.Register)

	//更新
	r.PUT("/user/profile") //鉴权中间件,controller.Update

	//发送邮箱验证码
	r.GET("/user/verify-code", controller.SendVerCodeByEmail)

	//重设密码(注册时发送“发送邮箱验证码”“重设密码”“注册”请求)
	r.GET("/users/verify-code", controller.ResetPassword)
}
