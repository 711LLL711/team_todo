package controller

//从路由请求中解析参数和结构体，调用 Model 和 Service，处理业务逻辑，决定如何响应用户的请求，处理异常和错误

import (
	"net/http"
	"team_todo/model"
	"team_todo/service"
	"team_todo/util"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

// 注册
func Register(c *gin.Context) {
	var userReq model.User
	userReq.Email = c.PostForm("email")
	userReq.Password = c.PostForm("password")
	userReq.Nickname = c.PostForm("nickname")
	userReq.Id = uuid.New().String()[:8]
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

	err := service.Modify(userReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok"})
}

// 发送邮箱验证码
func SendVerCodeByEmail(c *gin.Context) {
	service.SenderEmail()
	reqemail := c.PostForm("email")
	//检查邮箱是否合法
	if !util.IsValidEmail(reqemail) {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  "invalid email",
		})
		return
	}
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
