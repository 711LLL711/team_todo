package database

import (
	"log"
	"team_todo/config"
	"team_todo/global"
	"team_todo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type Database_Info struct {
// 	Username  string //数据库用户名
// 	Password  string
// 	Name      string //数据库名
// 	Hostname  string
// 	Port      string
// 	Parameter string //连接参数
// }

// 新建表
func Init() {
	// 创建用户表
	err := global.GVA_DB.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to create user table")
	}

	// 创建群组表
	err = global.GVA_DB.AutoMigrate(&model.Group{})
	if err != nil {
		panic("failed to create group table")
	}

	// 创建任务表
	err = global.GVA_DB.AutoMigrate(&model.Task{})
	if err != nil {
		panic("failed to create task table")
	}

	// 创建提醒表
	err = global.GVA_DB.AutoMigrate(&model.Reminder{})
	if err != nil {
		panic("failed to create reminder table")
	}
}

// 包括 数据库连接、查询数据库等函数
func Connect() {
	dsn := config.DSN(global.GVA_CONFIG.Database)
	var err error
	global.GVA_DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //连接数据库
	if err != nil {
		log.Fatal(err)
	}
}

// 建表
func CreateTables() {
	global.GVA_DB.Migrator().CreateTable(&model.User{})
	global.GVA_DB.Migrator().CreateTable(&model.VerCode{})
}
