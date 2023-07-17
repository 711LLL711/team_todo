package controller

//从路由请求中解析参数和结构体，调用 Model 和 Service，处理业务逻辑，决定如何响应用户的请求，处理异常和错误

import (
	"team_todo/model"
	"team_todo/service"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context){
	var userReq model.User
	userReq.Email = c.Postform("email")
	userReq.Password = c.Postform("password")
	userReq.Nickname = c.Postform("nickname")

	service.Register(userReq)
	
}

func Update(c *gin.Context){
	var userReq model.User
	userReq.Nickname = c.Postform("nickname")
	userReq.Avatar = c.Postform("avatar")

	service.Modify(userReq)
}