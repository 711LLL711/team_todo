package global

import (
	"log"
	"team_todo/config"

	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG *config.Config
	GVA_CWD    string
)

const BaseUrl = "http://localhost:8080"

// 解析配置文件到GVA_CONFIG
func LoadConfig() {
	configFilePath := "config/config.json"
	var err error
	GVA_CONFIG, err = config.LoadConfig(configFilePath)
	if err != nil {
		log.Println("解析失败")
		log.Fatal(err)
	}
	log.Println("解析成功")
}
