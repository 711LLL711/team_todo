package route

import (
	"team_todo/api/v1/controller"

	"github.com/gin-gonic/gin"
)

//设置路由
//登录、注册、找回密码...见api文档

//注册和更新
func MyRouter (r *gin.Engine){
	//注册
r.POST("/user/register",controller.Register)
	//更新
r.PUT("/user/profile",//鉴权中间件,controller.Update
)
}
