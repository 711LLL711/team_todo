package model

// 用于设置提醒的结构体
type Reminder struct {
	ReminderId string `json:"reminder"`
	TaskId     string `json:"taskId"`
	Deadline   string `json:"deadline"`
	Assignee   string `json:"assignee"`
	Nickname   string `json:"name"`
	Email      string `json:"emial"`
	Time       int    `json:"time"`
}
