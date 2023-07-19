package model

// 定义任务相关结构体
type Task struct {
	ID          string `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	GroupId     string `gorm:"column:groupid"`
	Description string
	Status      string //任务状态0-TODO 1-DONE
	AssigneeId  string `gorm:"column:assigneeid"` //负责人
	DueDate     string `gorm:"column:duedate"`    //ddl
}
