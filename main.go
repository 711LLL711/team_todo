package main

import (
	route "team_todo/api/v1"
	"team_todo/database"
	"team_todo/global"

	"fmt"

	"github.com/gin-gonic/gin"
)

// 程序入口
// 包括启动server,load 前端文件，连接数据库...
func main() {
	//load config
	global.LoadConfig()
	//connect the database
	database.Connect()

	database.CreateTables()
fmt.Println("connected")
	//设置服务器
	r := &gin.Engine{}
	route.SetupRoutes(r)
	Server := gin.Default()
	//load the front-end file
	//Server.LoadHTMLGlob("templates/*")
	Server.Run(":8080")

}
