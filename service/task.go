package service

import (
	"log"
	"team_todo/database"
	"team_todo/model"
)

//任务相关函数
//新建任务，删除任务，设置任务提醒

// 创建任务
func CreateTasks(name, description, status, assignee, deadline, groupId string) (taskID string, err error) {
	taskID, err1 := database.CreateTasks(name, description, status, assignee, deadline, groupId)
	if err1 != nil {
		log.Panicln("error")
		return "", err1
	}
	return taskID, nil
}

// 获取任务列表
func GetTasksList(groupId string) (count int, tasks []model.Task, err error) {
	count, tasks, err = database.GetTasksList(groupId)
	if err != nil {
		log.Panicln("error")
		return 0, nil, err
	}
	return count, tasks, err
}

// 获取任务信息
func GetTasks(taskID string) (task model.Task, err error) {
	log.Println("taskID:",taskID)
	task, err = database.GetTasks(taskID)
	log.Println("service task:",task)
	if err != nil {
		log.Panicln("error")
		return model.Task{}, err
	}
	return task, nil
}

// 更新任务信息
func ModifyTasks(taskID string, task model.Task) (groupId string, err error) {
	task,err = database.ModifyTasks(taskID, task)
	log.Println("service groupId:",task.GroupId)
	if err != nil {
		log.Panicln("error")
		return "", err
	}
	return task.GroupId, nil

}

// 删除任务
func DeleteTasks(taskID string) error {
	err := database.DeleteTasks(taskID)
	if err != nil {
		log.Panicln("error")
		return err
	}
	return nil
}
