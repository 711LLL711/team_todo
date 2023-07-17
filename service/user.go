package service

import (
	"team_todo/model"
	"team_todo/database"
)
//不确定要不要在这里返回错误
// 用户注册
func Register(userinfo model.User) {
	// 在这里调用数据库包中的函数来进行用户注册逻辑处理


	database.Register(userinfo)
	return
}

// 用户登录
func Login(userinfo model.User) {
	// 在这里调用数据库包中的函数来进行用户登录逻辑处理
	database.Login(userinfo)
	return
}

// 用户信息修改
func Modify(userinfo model.User) {
	// 在这里调用数据库包中的函数来进行用户信息修改逻辑处理
	database.Modify(userinfo)
	return 
}
