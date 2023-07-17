package team_todo

import (
	"team_todo/database"
	"team_todo/global"

	"github.com/gin-gonic/gin"
)

// 程序入口
// 包括启动server,load 前端文件，连接数据库...
func main() {
	//load config
	global.LoadConfig()
	//connect the database
	database.Connect()

	//设置服务器
	Server := gin.Default()
	//load the front-end file
	//Server.LoadHTMLGlob("templates/*")
	Server.Run(":8080")

}
