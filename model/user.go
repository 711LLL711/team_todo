package model

//定义用户结构体
type User struct {
	Nickname string
	Email    string
	Password string
	GroupId  []int  //用户所在的群组id
	avatar   string //头像图片存储在服务器的路径
}

//显示在个人主页的信息，只包括nickname和avatar
