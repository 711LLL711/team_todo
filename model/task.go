package model

// 定义任务相关结构体
type Task struct {
	ID          string `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	GroupId     string `gorm:"column:groupid" `
	Description string `json:"description"`
	Status      string `json:"status"`   //任务状态0-TODO 1-DONE
	Assignee    string `json:"assignee"` //负责人
	Deadline    string `json:"deadline"` //ddl
}
