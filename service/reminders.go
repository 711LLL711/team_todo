package service

import (
	"log"
	"team_todo/database"
	"team_todo/model"
)

// 创建提醒
func CreateReminders(taskId string,time int) (reminder model.Reminder, err error) {
	reminder, err = database.CreateReminders(taskId,time)
	if err != nil {
		log.Println("error:", err)
		return model.Reminder{}, err
	}
	return reminder, nil
}

// 获取提醒列表
func GetRemindersList(taskId string) (count int, reminders []model.Reminder, err error) {
	count, reminders, err = database.GetRemindersList(taskId)
	if err != nil {
		log.Panicln("error")
		return 0, nil, err
	}
	return count, reminders, err
}
// 删除提醒
func DeleteReminders(reminderId string) error {
	err := database.DeleteReminders(reminderId)
	if err != nil {
		log.Println("error:", err)
		return err
	}
	return nil
}
