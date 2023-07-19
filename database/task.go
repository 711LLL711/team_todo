package database

import (
	"team_todo/global"
	"team_todo/model"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// 创建任务
func CreateTasks(name, description, status, assignee, deadline, groupId string) (taskID string, err error) {
	var task model.Task
	task.Name = name
	task.Description = description
	task.Status = status
	task.Assignee = assignee
	task.Deadline = deadline
	task.GroupId = groupId
	// Generate a new UUID
	uuid, err1 := uuid.NewRandom()
	if err1 != nil {
		return "", errors.Wrap(err, "failed to generate UUID")
	}
	task.ID = uuid.String()
	result := global.GVA_DB.Create(&task)
	if result.Error != nil {
		return "", result.Error
	}
	return task.ID, nil
}

// 获取任务列表
func GetTasksList(groupId string) (count int, tasks []model.Task, err error) {
	result := global.GVA_DB.Where("group_id = ?", groupId).Find(&tasks)
	if result.Error != nil {
		return 0, nil, result.Error
	}
	// 获取数据条数
	count = len(tasks)
	return count, tasks, nil
}

// 获取任务信息
func GetTasks(taskID string) (task model.Task, err error) {
	result := global.GVA_DB.Where("ID = ?", taskID).Find(&task)
	if result.Error != nil {
		return model.Task{}, result.Error
	}
	return task, nil
}

// 更新任务信息
func ModifyTasks(taskID string, task model.Task) ( err error) {

	err = global.GVA_DB.Model(&model.Task{}).Where("id = ?", taskID).Updates(&task).Error
	if err != nil{
		return err
	}
	return nil
}

//删除任务
func DeleteTasks(taskID string) error{
	result := global.GVA_DB.Where("ID = ?", taskID).Delete(&model.Task{})
	
	if result.Error != nil {
		return result.Error
	}

	return nil
}

