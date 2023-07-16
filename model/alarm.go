package model

import (
	"time"
)

// 用于设置提醒的结构体
type alarm struct {
	TaskId  int
	DueDate time.Time
}
