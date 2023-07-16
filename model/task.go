package model

import (
	"time"
)

// 定义任务相关结构体
type Task struct {
	ID          int    `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Status      string //任务状态0-TODO 1-DONE
	Assignee    string //负责人
	DueDate     time.Time
}
