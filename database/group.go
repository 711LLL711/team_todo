package database

import (
	"fmt"
	"strings"
	"team_todo/global"
	"team_todo/model"

	"gorm.io/gorm"
)

func CreateGroup(groupinfo model.Group) error {
	result := global.GVA_DB.Create(&groupinfo)
	if result.Error != nil {
		// 处理错误，例如打印日志或返回自定义错误消息
		return result.Error
	}
	// 如果没有错误，返回 nil 表示操作成功
	return nil
}
func GetUserGroups(userId string) ([]model.ShowGroup, int, error) {
	var groupstr string

	if err := global.GVA_DB.Table("users").Where("id = ?", userId).First(&groupstr).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, 0, fmt.Errorf("user not found")
		}
		return nil, 0, err
	}

	var groups []model.ShowGroup
	if len(groupstr) > 0 {
		groupIds := strings.Split(groupstr, ",")
		if err := global.GVA_DB.Find(&groups, "group_id IN ?", groupIds).Error; err != nil {
			return nil, 0, err
		}
	}

	return groups, len(groups), nil
}

func GetGroupInfo(GroupId string) (model.Group, error) {
	var groupinfo model.Group
	err := global.GVA_DB.Table("groups").Where("group_id = ?", GroupId).First(&groupinfo).Error
	if err != nil {
		return groupinfo, err
	}
	return groupinfo, nil
}

func GetGroupMembers(GroupId string) ([]model.User, int64, error) {
	var users []model.User
	var count int64
	if err := global.GVA_DB.Select("nickname, id, avatar, email").
		Where("group_id LIKE ?", "%"+GroupId+"%").
		Find(&users).Error; err != nil {
		return nil, 0, err
	}
	if err := global.GVA_DB.Model(&model.User{}).
		Where("group_id LIKE ?", "%"+GroupId+"%").
		Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return users, count, nil

}
