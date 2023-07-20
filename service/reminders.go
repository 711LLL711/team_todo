package service

import (
	"log"
	"team_todo/database"
	"team_todo/model"
)

// 创建提醒
func CreateReminders(taskId string) (reminder model.Reminder, err error) {
	reminder, err = database.CreateReminders(taskId)
	if err != nil {
		log.Println("error:", err)
		return model.Reminder{}, err
	}
	return reminder, nil
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
