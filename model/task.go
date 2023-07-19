package model

// 定义任务相关结构体
type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Assignee    string `json:"assignee"`
	Deadline    string `json:"deadline"`
	GroupId     string `json:"groupId"`
}
