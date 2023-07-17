package service

import (
	"team_todo/database"
	"team_todo/model"
)

//群组相关函数
//创建群组，邀请入群

// 创建群组
func CreateGroup(groupinfo model.Group) {
	database.CreateGroup(groupinfo)
	return
}
