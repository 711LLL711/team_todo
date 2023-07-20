package controller

import (
	"log"
	"net/http"
	"team_todo/model"
	"team_todo/service"

	"github.com/gin-gonic/gin"
)

// 创建任务
func CreateTasks(c *gin.Context) {
	var task model.Task
	task.Name = c.PostForm("name")
	task.Description = c.PostForm("description")
	task.Status = c.PostForm("status")
	task.Assignee = c.PostForm("assignee")
	task.Deadline = c.PostForm("deadline")
	task.GroupId = c.PostForm("groupId")

	var err1 error
	task.ID, err1 = service.CreateTasks(task.Name, task.Description, task.Status, task.Assignee, task.Deadline, task.GroupId)
	if err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "wrong"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// 获取任务列表
func GetTasksList(c *gin.Context) {
	var tasks []model.Task
	var groupId string
	groupId = c.Query("groupId")
	var count int
	var err error
	count, tasks, err = service.GetTasksList(groupId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "wrong"})
		return
	}

	var response []struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Status      string `json:"status"`
		Assignee    string `json:"assignee"`
		Deadline    string `json:"deadline"`
		GroupID     string `json:"groupId"`
	}

	for _, task := range tasks {
		response = append(response, struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
			Status      string `json:"status"`
			Assignee    string `json:"assignee"`
			Deadline    string `json:"deadline"`
			GroupID     string `json:"groupId"`
		}{
			ID:          task.ID,
			Name:        task.Name,
			Description: task.Description,
			Status:      task.Status,
			Assignee:    task.Assignee,
			Deadline:    task.Deadline,
			GroupID:     task.GroupId,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"count": count,
		"task":  response,
	})
}

// 获取任务信息
func GetTasks(c *gin.Context) {
	id := c.Param("id")
	var task model.Task
	var err error
	task, err = service.GetTasks(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "wrong"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// 更新任务信息
func ModifyTasks(c *gin.Context) {
	taskID := c.Param("id")
	var task model.Task
	task.ID = c.PostForm("id")
	if taskID !=task.ID{
		c.JSON(http.StatusBadRequest, gin.H{"err": "任务ID冲突"})
		log.Panicln("error")
		return
	}
	task.Name = c.PostForm("name")
	task.Description = c.PostForm("description")
	task.Status = c.PostForm("status")
	task.GroupId = c.PostForm("groupId")

	var err error
	task.GroupId, err = service.ModifyTasks(taskID, task)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": "wrong"})
		return
	}
	c.JSON(http.StatusOK, task)
}

//删除任务
func DeleteTasks(c *gin.Context) {
	taskID := c.Param("id")
	err := service.DeleteTasks(taskID)
	
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"err": "wrong"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"success"})
}
