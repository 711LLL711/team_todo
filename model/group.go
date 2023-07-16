package model

//群组相关结构体，函数
type Group struct {
	GroupId         int
	GroupName       string
	Group_Member_Id []int
	Group_Task_Id   []int
}
