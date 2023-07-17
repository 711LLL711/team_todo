package model

import (
	"time"
)

// 用于设置提醒的结构体
type Reminder struct {
	ReminderId int
	TaskId     int
	DueDate    time.Time
}
