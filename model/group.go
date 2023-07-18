package model

//群组相关结构体
type Group struct {
	GroupId           string
	GroupName         string
	GroupOwnerId      string
	Group_Description string
	Group_Invite_Id   string //邀请码
}
type ShowGroup struct {
	GroupId           string
	GroupName         string
	Group_Description string
}
type QueryGroup struct {
	GroupId           string
	GroupName         string
	Group_Description string
	GroupOwnerId      string
}

//群组成员id对应表
type GroupWithUser struct {
	UserId  string
	GroupId string
}
