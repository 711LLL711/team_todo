package controller

//从路由请求中解析参数和结构体，调用 Model 和 Service，处理业务逻辑，决定如何响应用户的请求，处理异常和错误

//处理用户请求相关逻辑

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"team_todo/database"
	"team_todo/global"
	"team_todo/model"
	"team_todo/service"
	"team_todo/util"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	var loginReq model.User
	loginReq.Email = c.PostForm("email")
	loginReq.Password = c.PostForm("password")
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

// 用户上传头像图片
func (uc *UserController) UploadAvatar(c *gin.Context) {
	// Parse incoming image file
	file, err := c.FormFile("image")
	if err != nil {
		log.Println("image upload error --> ", err)
		c.JSON(500, gin.H{"status": 500, "message": "Server error", "data": nil})
		return
	}

	// Generate a new UUID for the image name
	uniqueID := uuid.New()

	// Remove "-" from the image name
	filename := strings.Replace(uniqueID.String(), "-", "", -1)

	// Extract the image extension from the original file filename
	fileExt := strings.Split(file.Filename, ".")[1]

	// Generate the image name with the filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)
	// Save the image to the "./images" directory
	err = c.SaveUploadedFile(file, fmt.Sprintf("./images/%s", image))
	if err != nil {
		log.Println("image save error --> ", err)
		c.JSON(500, gin.H{"status": 500, "message": "Server error", "data": nil})
		return
	}

	// Generate the image URL to serve to the client using CDN
	imageURL := fmt.Sprintf("%s/images/%s", global.BaseUrl, image)
	//保存用户头像到数据库
	userid, _ := util.GetId_Session(c.Request)
	err = database.UpdateAvatar(userid, imageURL)
	if err != nil {
		c.JSON(500, gin.H{"status": 500, "message": "Server error", "data": nil})
		return
	}
	c.JSON(http.StatusOK, gin.H{"url": imageURL})
}
