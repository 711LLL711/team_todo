package controller

//从路由请求中解析参数和结构体，调用 Model 和 Service，处理业务逻辑，决定如何响应用户的请求，处理异常和错误

//处理用户请求相关逻辑

import (
	"net/http"

	"team_todo/model"
	"team_todo/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	// 可以在这里定义其他依赖的服务或存储库
}

// 登录页面
func (uc *UserController) ShowLoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "sign_&_log.html", gin.H{
		"title": "登录",
	})
}

// 处理登录请求
func (uc *UserController) Login(c *gin.Context) {
	var loginReq model.LoginReq
	if err := c.ShouldBind(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用服务层处理登录逻辑
	err := service.Login(loginReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
}
