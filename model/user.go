package model

//定义用户结构体
type User struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"` //头像图片存储在服务器的路径
}
type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type QueryUser struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
}
