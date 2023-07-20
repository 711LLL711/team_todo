package database

import (
	"errors"
	"log"
	"team_todo/global"
	"team_todo/model"

	"gorm.io/gorm"
)

func CreateGroup(groupinfo model.Group) error {
	result := global.GVA_DB.Table("group").Create(&groupinfo)
	if result.Error != nil {
		// 处理错误，例如打印日志或返回自定义错误消息
		return result.Error
	}
	// 如果没有错误，返回 nil 表示操作成功
	return nil
}

// 添加用户到群组
func AddUserToGroup(groupwithuser model.GroupWithUser) error {
	result := global.GVA_DB.Table("groupwithuser").Create(&groupwithuser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 查询用户加入的所有群组
func GetUserGroups(UserId string) ([]model.ShowGroup, int64, error) {
	var groupIDs []string
	if err := global.GVA_DB.Table("groupwithuser").
		Select("groupid").
		Where("userid = ?", UserId).
		Find(&groupIDs).Error; err != nil {
		return nil, 0, err
	}

	// 统计 group 数量
	count := int64(len(groupIDs))

	// 查询 group 表中对应 groupid 的信息
	var showGroups []model.ShowGroup
	if err := global.GVA_DB.Table("group").
		Select("groupid, groupname, group_description").
		Where("groupid IN ?", groupIDs).
		Find(&showGroups).Error; err != nil {
		return nil, 0, err
	}
	log.Println("showGroups: ", showGroups)
	return showGroups, count, nil
}

// 查询群组信息
func GetGroupInfo(GroupId string) (model.Group, error) {
	var groupinfo model.Group
	err := global.GVA_DB.Table("group").Where("groupid = ?", GroupId).First(&groupinfo).Error
	if err != nil {
		return groupinfo, err
	}
	return groupinfo, nil
}

func GetGroupMembers(GroupId string) ([]model.QueryUser, int64, error) {
	var userIDs []string
	var users []model.QueryUser
	var count int64

	// 查询 groupwithuser 表获取 userid 列表
	if err := global.GVA_DB.Table("groupwithuser").
		Select("userid").
		Where("groupid = ?", GroupId).
		Find(&userIDs).Error; err != nil {
		return nil, 0, err
	}

	// 查询 user 表获取对应的记录和总记录数
	if err := global.GVA_DB.Table("user").
		Select("id, nickname, email, avatar").
		Where("id IN ?", userIDs).
		Find(&users).Error; err != nil {
		return nil, 0, err
	}

	// 查询 userid 总数
	if err := global.GVA_DB.Table("groupwithuser").
		Where("groupid = ?", GroupId).
		Count(&count).Error; err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func GetGroupCode(GroupId string) (string, error) {
	var group model.Group
	if err := global.GVA_DB.Table("group").
		Select("group_invite_id").
		First(&group, "groupid = ?", GroupId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("group not exist")
		}
	}

	return group.Group_Invite_Id, nil
}
func StoreGroupCode(GroupId string, invitecode string) error {
	var group model.Group

	if err := global.GVA_DB.Table("group").Where("groupid = ?", GroupId).First(&group).Error; err != nil {
		return err
	}
	group.Group_Invite_Id = invitecode
	if err := global.GVA_DB.Table("group").Where("groupid = ?", GroupId).Updates(&group).Error; err != nil {
		return err
	}

	return nil
}

// 根据inviteid查找对应的群组id
func Search_Group_Invite(code string) (string, error) {
	var GroupId string

	// 查询指定的 group_invite_id 记录
	if err := global.GVA_DB.Table("group").
		Where("group_invite_id = ?", code).
		Pluck("groupid", &GroupId).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 没有找到匹配的记录
			return "", errors.New("NOT FOUND")
		}
		// 查询过程中发生错误
		return "", err
	}

	// 找到匹配的记录
	return GroupId, nil
}

// 退出群组
func LeaveGroup(groupwithuser model.GroupWithUser) error {
	if err := global.GVA_DB.Table("groupwithuser").
		Where("userid = ? AND groupid = ?", groupwithuser.UserId, groupwithuser.GroupId).
		Delete(&model.GroupWithUser{}).Error; err != nil {
		return err
	}

	return nil
}

func DelGroup(GroupId string) error {
	// 开启事务
	tx := global.GVA_DB.Begin()

	// 删除 groupwithuser 表中的记录
	if err := tx.Table("groupwithuser").
		Where("groupid = ?", GroupId).
		Delete(&model.GroupWithUser{}).Error; err != nil {
		// 发生错误时回滚事务
		tx.Rollback()
		return err
	}

	// 删除 group 表中的记录
	if err := tx.Table("group").
		Where("groupid = ?", GroupId).
		Delete(&model.Group{}).Error; err != nil {
		// 发生错误时回滚事务
		tx.Rollback()
		return err
	}

	// 查询属于指定 groupID 的 taskid
	var taskIDs []string
	if err := tx.Table("task").
		Where("groupid = ?", GroupId).
		Pluck("id", &taskIDs).Error; err != nil {
		// 发生错误时回滚事务
		tx.Rollback()
		return err
	}

	// 删除属于指定 groupID 的 taskid 在 reminder 表中的记录
	if err := tx.Table("reminder").
		Where("taskid IN ?", taskIDs).
		Delete(&model.Reminder{}).Error; err != nil {
		// 发生错误时回滚事务
		tx.Rollback()
		return err
	}

	// 删除属于指定 groupID 的 task 表中的记录
	if err := tx.Table("task").
		Where("groupid = ?", GroupId).
		Delete(&model.Task{}).Error; err != nil {
		// 发生错误时回滚事务
		tx.Rollback()
		return err
	}

	// 提交事务
	tx.Commit()

	return nil
}

func IsGroupOwner(GroupId, UserId string) bool {
	var group model.GroupWithUser
	if err := global.GVA_DB.Table("groupwithuser").
		Where("groupid = ?", GroupId).
		First(&group).Error; err != nil {
		return false
	}
	if group.UserId == UserId {
		return true
	}
	return false
}
