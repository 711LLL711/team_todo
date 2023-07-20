package database

import (
	"team_todo/global"
	"team_todo/model"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// 从数据库里查询reminder表
func RemindNow() (reminders []model.Reminder, err error) {
	if err = global.GVA_DB.Table("reminder").Find(&reminders).Error; err != nil {
		return nil, err
	}

	return reminders, nil
}

func CreateReminders(taskId string) (reminder model.Reminder, err error) {
	// Generate a new UUID
	uuid, err1 := uuid.NewRandom()
	if err1 != nil {
		return model.Reminder{}, errors.Wrap(err, "failed to generate UUID")
	}
	reminder.ReminderId = uuid.String()
	reminder.TaskId = taskId
	task, err1 := GetTasks(taskId)
	if err1 != nil {
		return model.Reminder{}, err1
	}

	reminder.Assignee, reminder.Deadline = task.Assignee, task.Deadline
	user, err2 := GetProfile(reminder.Assignee)
	if err2 != nil {
		return model.Reminder{}, err2
	}
	reminder.Email, reminder.Nickname = user.Email, user.Nickname
	result := global.GVA_DB.Create(&reminder)
	if result.Error != nil {
		return model.Reminder{}, result.Error
	}
	return reminder, nil
}

// 删除提醒
func DeleteReminders(reminderId string) error {

	result := global.GVA_DB.Where("reminder_id = ?", reminderId).Delete(&model.Reminder{})

	if result.Error != nil {
		return result.Error
	}

	return nil

}
