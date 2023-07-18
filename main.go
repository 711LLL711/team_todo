package main

import (
	"team_todo/database"
	"team_todo/global"
	//"team_todo/model"
)

// 程序入口
// 包括启动server,load 前端文件，连接数据库...
func main() {
	//load config
	global.LoadConfig()
	database.Connect()
	database.CreateTables()
}
