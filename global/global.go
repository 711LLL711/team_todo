package global

import (
	"log"
	"team_todo/config"

	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG *config.Config
)

// 解析配置文件到GVA_CONFIG
func LoadConfig() {
	GVA_CONFIG = &config.Config{}
	configFilePath := "config/config.json"
	var err error
	*GVA_CONFIG, err = config.LoadConfig(configFilePath)
	if err != nil {
		log.Fatal(err)
	}
}
