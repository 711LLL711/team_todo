package route

import (
	"team_todo/api/v1/controller"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine) {
	// 创建任务
	r.POST("/task",
		//jwt等身份处理
		controller.CreateTasks,
	)
	// 获取任务列表
	r.GET("/tasks",
		//jwt等身份处理
		controller.GetTasksList)
	// 获取任务信息
	r.GET("/tasks/:id",
		//jwt等身份处理
		controller.GetTasks,
	)
	// 更新任务信息
	r.PUT("/tasks/:id",
		//jwt等身份处理
		controller.ModifyTasks)
	//删除任务
	r.DELETE("/tasks/:id",
		//jwt等身份处理
		controller.DeleteTasks)
}
