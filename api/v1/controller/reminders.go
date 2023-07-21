package controller

import (
	"log"
	"net/http"
	"team_todo/model"
	"team_todo/service"

	"github.com/gin-gonic/gin"
)

// 创建提醒
func CreateReminders(c *gin.Context) {
	var reminder model.Reminder
	var err error
	reminder.TaskId = c.PostForm("taskId")
	reminder, err = service.CreateReminders(reminder.TaskId)
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
