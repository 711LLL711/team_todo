package database

import (
	"log"
	"team_todo/config"
	"team_todo/global"

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

// 包括 数据库连接、查询数据库等函数
func Connect() {
	dsn := config.DSN(global.GVA_CONFIG.Database)
	var err error
	global.GVA_DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}) //连接数据库
	if err != nil {
		log.Fatal(err)
	}
}
