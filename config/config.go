package config

//解析配置文件
import (
	"encoding/json"
	"os"
)

type Config struct {
	JWT_secret string
	Database   DatabaseConfig
	Email EmailConfig
}

type DatabaseConfig struct {
	Username  string //数据库用户名
	Password  string
	Name      string //数据库名
	Hostname  string
	Port      string
	Parameter string //连接参数
}

type EmailConfig struct{
	SMTP_server string
	SenderEmail string 
	SenderPassword string
}
func LoadConfig(filePath string) (Config, error) {
	var config Config

	file, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// DSN returns the Data Source Name
func DSN(ci DatabaseConfig) string {
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
