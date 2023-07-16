package database

import (
	"database/sql"
)

type Database_Info struct {
	Username  string //数据库用户名
	Password  string
	Name      string //数据库名
	Hostname  string
	Port      string
	Parameter string //连接参数
}

var Db *sql.DB

// DSN returns the Data Source Name
func DSN(ci Database_Info) string {
	// Example: root:@tcp(localhost:3306)/test
	return ci.Username +
		":" +
		ci.Password +
		"@tcp(" +
		ci.Hostname +
		":" +
		ci.Port +
		")/" +
		ci.Name + ci.Parameter
}

// 包括 数据库连接、查询数据库等函数
func Connect() //待实现--不带参数，直接load配置文件连接
