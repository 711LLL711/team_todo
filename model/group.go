package model

//群组相关结构体
type Group struct {
	GroupId           string `gorm:"column:groupid"`
	GroupName         string `gorm:"column:groupname"`
	GroupOwnerId      string `gorm:"column:groupownerid"`
	Group_Description string `gorm:"column:group_description"`
	Group_Invite_Id   string `gorm:"column:group_invite_id"` //邀请码
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
	UserId  string `gorm:"column:userid"`
	GroupId string `gorm:"column:groupid"`
}
