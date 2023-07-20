package main

import (
	route "team_todo/api/v1"
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
	//建表
	database.CreateTables()
	//设置服务器
	Server := gin.Default()

	//设置存放图片
	Server.Static("/images", "./images")

	//设置存放静态文件
	Server.Static("/static", "./static")
	Server.Static("/javascripts", "./templates/javascripts")
	Server.Static("/stylesheets", "./templates/stylesheets")

	//load html
	Server.LoadHTMLGlob("templates/views/*")
	// 设置路由
	route.SetupRoutes(Server)

	route.TaskRoutes(Server)
	//load the front-end file
	//Server.LoadHTMLGlob("templates/*")
	Server.Run(":8080")

}
