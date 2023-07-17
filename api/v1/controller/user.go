package controller

//从路由请求中解析参数和结构体，调用 Model 和 Service，处理业务逻辑，决定如何响应用户的请求，处理异常和错误

//处理用户请求相关逻辑

import (
	"net/http"
	"team_todo/model"
	"team_todo/service"
	"team_todo/util"

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

	// 调用服务层验证密码，生成session
	err := service.Login(loginReq, c.Request, c.Writer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//设置token并返回authorization头部
	token, expireTimestap, err := util.GenerateToken(loginReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate JWT"})
		return
	}
	c.Header("Authorization", "Bearer "+token)
	c.JSON(http.StatusOK, gin.H{"token": token, "expire": expireTimestap})
}

func (uc *UserController) GetProfile(c *gin.Context) {
	GetId := c.Param("id")
	user, err := service.GetProfile(GetId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": user.Id, "nickname": user.Nickname, "avatar": user.Avatar})
}
