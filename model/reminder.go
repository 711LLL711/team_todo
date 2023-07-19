package model

// 用于设置提醒的结构体
type Reminder struct {
	ReminderId string `gorm:"column:reminderid"`
	TaskId     string `gorm:"column:taskid"`
	DueDate    string `gorm:"column:duedate"`
}
