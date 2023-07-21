package database

import (
	"log"
	"team_todo/global"
	"team_todo/model"
	"time"

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

func RemindNow1() (reminders []model.Reminder, err error) {
	//查询预定时间在这一分钟的函数
	min :=time.Now().Unix()
	max :=time.Now().Add(time.Minute).Unix()
	if err = global.GVA_DB.Table("reminder").Where("time >= ? AND time < ?", min, max).Find(&reminders).Error; err != nil {
		return nil, err
	}

	return reminders, nil
}
//创建提醒
func CreateReminders(taskId string,time int) (reminder model.Reminder, err error) {
	// Generate a new UUID
	uuid, err1 := uuid.NewRandom()
	if err1 != nil {
		return model.Reminder{}, errors.Wrap(err, "failed to generate UUID")
	}
	reminder.ReminderId = uuid.String()
	reminder.TaskId = taskId
	reminder.Time = time
	task, err1 := GetTasks(taskId)
	log.Println("database create a reminder GetTask() taskId:",taskId)
	if err1 != nil {
		return model.Reminder{}, err1
	}
log.Println("database create a reminder search for task: ",task)
	reminder.Assignee, reminder.Deadline = task.Assignee, task.Deadline
	user, err2 := GetProfile(reminder.Assignee)
	if err2 != nil {
		return model.Reminder{}, err2
	}
	log.Println("database create a reminder search for user:",user)
	reminder.Email, reminder.Nickname = user.Email, user.Nickname
	result := global.GVA_DB.Table("reminder").Create(&reminder)
	if result.Error != nil {
		return model.Reminder{}, result.Error
	}
	return reminder, nil
}
// 获取提醒列表
func GetRemindersList(taskId string) (count int, reminders []model.Reminder, err error) {
	result := global.GVA_DB.Table("reminder").Where("task_id = ?", taskId).Find(&reminders)
	if result.Error != nil {
		return 0, nil, result.Error
	}
	// 获取数据条数
	count = len(reminders)
	return count, reminders, nil
}
// 删除提醒
func DeleteReminders(reminderId string) error {

	result := global.GVA_DB.Table("reminder").Where("reminder_id = ?", reminderId).Delete(&model.Reminder{})

	if result.Error != nil {
		return result.Error
	}

	return nil

}
