package controller

import (
	"log"
	"net/http"
	"strconv"
	"team_todo/model"
	"team_todo/service"

	"github.com/gin-gonic/gin"
)

// 创建提醒
func CreateReminders(c *gin.Context) {
	var reminder model.Reminder
	var err error
	reminder.TaskId = c.PostForm("taskId")
timeStr:= c.PostForm("time")
timeInt, err := strconv.Atoi(timeStr)
		if err != nil {
			c.String(400, "Invalid time value")
			return
		}
	reminder.Time  = timeInt
	log.Println("reminder.Time:",reminder.Time)
	reminder, err = service.CreateReminders(reminder.TaskId,reminder.Time)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, reminder)
}

//获取提醒列表
// func GetReminderList(c *gin.Context){

// }

func GetRemindersList(c *gin.Context) {
	var reminders []model.Reminder
	taskId := c.Query("taskId")
	var count int
	var err error
	count, reminders, err = service.GetRemindersList(taskId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "wrong"})
		return
	}

	var response []model.Reminder

	for _, reminder := range reminders {
		response = append(response,model.Reminder{
				ReminderId: reminder.ReminderId,
				TaskId:     reminder.TaskId,
				Deadline:   reminder.Deadline,
				Assignee:   reminder.Assignee,
				Nickname:   reminder.Nickname,
				Email:      reminder.Email,
				Time:       reminder.Time,
			})
	}

	c.JSON(http.StatusOK, gin.H{
		"count": count,
		"task":  response,
	})
}

// 删除提醒
func DeleteReminders(c *gin.Context) {
	reminderId := c.Param("id")
	log.Println("controller deletereminders id:",reminderId)
	err := service.DeleteReminders(reminderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
