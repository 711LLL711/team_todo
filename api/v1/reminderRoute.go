package route

import (
	"team_todo/api/v1/controller"

	"github.com/gin-gonic/gin"
)

func RouteForReminder(r *gin.Engine) {
	//创建提醒
	r.POST("/reminders", controller.CreateReminders)
	//获取提醒列表
	r.GET("/reminders", controller.GetRemindersList)
	//删除提醒
	r.DELETE("/reminders/:id", controller.DeleteReminders)
}
