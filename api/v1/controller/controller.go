package controller

//从路由请求中解析参数和结构体，调用 Model 和 Service，处理业务逻辑，决定如何响应用户的请求，处理异常和错误

import (
	"team_todo/model"
	"team_todo/service"

	"net/http"

	"github.com/gin-gonic/gin"
)

// 注册
func Register(c *gin.Context) {
	var userReq model.User
	userReq.Email = c.PostForm("email")
	userReq.Password = c.PostForm("password")
	userReq.Nickname = c.PostForm("nickname")

	if userReq.Email == "" || userReq.Password == "" || userReq.Nickname == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "信息为空"})
		return
	}

	err := service.Register(userReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": userReq.Password,
		"email":    userReq.Email,
		"nickname": userReq.Nickname})

}

// 更新信息
func Update(c *gin.Context) {
	var userReq model.User
	userReq.Nickname = c.PostForm("nickname")
	userReq.Avatar = c.PostForm("avatar")

	service.Modify(userReq)
}

// 发送邮箱验证码
func SendVerCodeByEmail(c *gin.Context) {
	service.SenderEmail()
	reqemail := c.PostForm("email")
	code := service.GenVerCode(reqemail) //生成验证码，并存到数据库
	err := service.PerformEmailSending(reqemail, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err,
		})
		return
	}
}

// 重设密码
func ResetPassword(c *gin.Context) {
	var userReq model.User
	userReq.Email = c.PostForm("email")
	userReq.Password = c.PostForm("password")
	code := c.PostForm("code")
	res := service.CheckVercode(code, userReq.Email)
	if res {
		service.Modify(userReq)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "fail",
		"error":  "wrong verification code",
	})
}
