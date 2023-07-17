package database

import (
	"fmt"
	"team_todo/model"
	"team_todo/global"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"crypto/sha256"
	"encoding/hex"

)

//加密
func Myencrypt(row string) string{
    data := []byte(row)
    hash := sha256.Sum256(data)
    fmt.Printf("%x\n", hash)
	res := hex.EncodeToString(hash[:])
	return res
}

// 由service调用的注册函数
func Register(userinfo model.User) {
	//DB是数据库连接

	exist := Exists(global.GVA_DB, "users", "email", userinfo.Email)

	// fmt.Println("whether the email is used: ", exist)
	if exist == true {
		// err1 := fmt.Errorf("y")
		// panic(err1)

		//返回已注册
		fmt.Println("The email is already registered.Change a email number or log in directly.")
		return
	} else {
		userinfo.Password = (userinfo.Password)
		if userinfo.Password != "" && userinfo.Nickname != "" {
			global.GVA_DB.Create(&user)
			fmt.Println("oh created")
		}
	}
}

// 判断指定表tableName的字段column是否存在值value
func Exists(db *gorm.DB, tableName string, column string, value interface{}) bool {

	// 使用map进行动态查询
	condition := map[string]interface{}{
		column: value,
	}

	var count int64
	db.Table(tableName).Where(condition).Count(&count)
	fmt.Println("exist number: ", count)
	return count > 0 //bool count == 0 false-->不存在；count > 0 true-->已存在
}

// 由service调用的登录函数
func Login(userinfo model.User) {
		var being_logged model.User
		global.GVA_DB.Table("users").Where("email = ?", userinfo.Email).First(&being_logged)

		
		switch {
		case Myencrypt(userinfo.Password) == being_logged.Password:
			fmt.Println("log in successfully")
			//登录成功标识......
			//......
		case being_logged.Email != "":
			fmt.Println("The email and the password is not matched.")
			//登录失败
		default:
			fmt.Println("The email number is not registered.Please register first.")
			//未注册
		}
}

// 由service调用的用户信息修改函数
func Modify(userinfo model.User,nickname string,avatar string) {
	var userReq model.User
	userReq.Nickname = nickname
	userReq.Avatar = avatar
	global.GVA_DB.Model(&model.User{}).Updates(userReq)
}
/*
//上传头像图片函数
func UploadAvatar(c *gin.Context) {
	// 从请求中获取上传的文件
	file, err := c.FormFile("avatar")
	if err != nil {
		// 处理错误
		c.JSON(400, gin.H{
			"error": "Failed to retrieve avatar file",
		})
		return
	}

	// 将文件保存到服务器端指定的目录或云存储服务中
	err = c.SaveUploadedFile(file, "/path/to/avatar_directory/"+file.Filename)
	if err != nil {
		// 处理错误
		c.JSON(500, gin.H{
			"error": "Failed to save avatar file",
		})
		return
	}

	// 返回上传后的图片链接
	avatarURL := "https://example.com/avatars/" + file.Filename
	c.JSON(200, gin.H{
		"url": avatarURL,
	})
}
*/
