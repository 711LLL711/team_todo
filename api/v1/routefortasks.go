package route

import (
	"net/http"
	"team_todo/api/v1/controller"

	"github.com/gin-gonic/gin"
)

func TaskRoutes(r *gin.Engine){
	r.POST("/task",
	//jwt等身份处理
	controller.CreateTasks,
)
r.GET("/tasks",
	//jwt等身份处理
controller.GetTasksList)
r.GET("/tasks/:id",
	//jwt等身份处理
controller.GetTasks, 
	)
r.PUT("/tasks",
	//jwt等身份处理
controller.ModifyTasks)
r.DELETE("/tasks",
	//jwt等身份处理
controller.DeleteTasks)
}