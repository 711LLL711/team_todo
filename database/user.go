package database

import (
	"fmt"
	"team_todo/model"

	"github.com/jinzhu/gorm"
)

// 由service调用的注册函数
func Register(userinfo model.User) {
	//DB是数据库连接
	exist := Exists(Db, "users", "email", userinfo.Email)

	// fmt.Println("whether the email is used: ", exist)
	if exist == true {
		// err1 := fmt.Errorf("y")
		// panic(err1)

		//返回已注册
		fmt.Println("The email is already registered.Change a email number or log in directly.")
		return
	} else {
		if userinfo.Password != "" && userinfo.Nickname != "" {
			Db.Create(&user)
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
		Db.Table("users").Where("email = ?", userinfo.Email).First(&being_logged)

		//err3 := db.Table("users").Where("phone_number = ?", user.PhoneNumber).First(&being_logged).Error

		//checkerr(err3)

		// fmt.Println("user password: ", user.Password, "user Nickname: ", user.NickName, "right password: ", being_logged.Password, "right nickname: ", being_logged.NickName)
		// fmt.Println("the user who is being logged : ", being_logged.PhoneNumber, being_logged.Password)
		// fmt.Println("user password: ", user.Password, "right password: ", being_logged.Password)
		switch {
		case userinfo.Password == being_logged.Password:
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
	userReq.avatar = avatar
	Db.Model(User).Updates(userReq)
}
