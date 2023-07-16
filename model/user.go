package model

//定义用户结构体
type User struct {
	Nickname string
	Email    string
	Password string
	GroupId  []int  //用户所在的群组id
	Avatar   string //头像图片存储在服务器的路径
}
type LoginReq struct {
	Email    string
	Password string
}

//显示在个人主页的信息，只包括nickname和avatar
