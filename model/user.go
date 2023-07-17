package model

//定义用户结构体
type User struct {
	Id       string
	Nickname string
	Email    string
	Password string
	GroupId  []string //用户所在的群组id
	Avatar   string   //头像图片存储在服务器的路径
}
type LoginReq struct {
	Email    string
	Password string
}
type QueryUser struct {
	Id       string
	Nickname string
	Email    string
	Avatar   string
}
