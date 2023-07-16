package team_todo

import (
	"team_todo/global"
)

// 程序入口
// 包括启动server,load 前端文件，连接数据库...
func main() {
	//load config
	global.LoadConfig()
}
