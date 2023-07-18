package model

// 定义任务相关结构体
type Task struct {
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	GroupId     string
	Description string
	Status      string //任务状态0-TODO 1-DONE
	AssigneeId  string //负责人
	DueDate     string //ddl
}
