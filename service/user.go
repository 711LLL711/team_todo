package service

import (
	"net/http"
	"team_todo/database"
	"team_todo/model"
	"team_todo/util"
)

// 不确定要不要在这里返回错误
// 用户注册
func Register(userinfo model.User) {
	// 在这里调用数据库包中的函数来进行用户注册逻辑处理

	database.Register(userinfo)
	return
}

// 用户登录
func Login(userinfo model.LoginReq, req *http.Request, resp http.ResponseWriter) error {
	// 在这里调用数据库包中的函数来进行用户登录逻辑处理
	err := database.Login(userinfo)
	if err != nil {
		return err
	}
	//设置session
	var session util.Session_info
	session.Email = userinfo.Email
	session.Id, err = database.GetId(userinfo.Email)
	if err != nil {
		return err
	}
	err = util.SetSession(resp, req, session.Id, session.Email)
	if err != nil {
		return err
	}
	return nil
}

// 用户信息修改
func Modify(userinfo model.User, nickname string, avatar string) {
	// 在这里调用数据库包中的函数来进行用户信息修改逻辑处理
	database.Modify(userinfo, nickname, avatar)
	return
}

//获取用户资料

func GetProfile(UserId string) (model.User, error) {
	user, err := database.GetProfile(UserId)
	return user, err
}
