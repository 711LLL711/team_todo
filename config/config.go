package config

//解析配置文件
import (
	"encoding/json"
	"os"
)

type Config struct {
	Database   DatabaseConfig `json:"database"`
	JWT_secret string         `json:"jwt_secret"`
	Email      EmailConfig    `json:"email"`
}

type DatabaseConfig struct {
	Hostname  string `json:"host"`
	Port      string `json:"port"`
	Username  string `json:"username"` // 数据库用户名
	Password  string `json:"password"`
	Name      string `json:"name"`      // 数据库名
	Parameter string `json:"parameter"` // 连接参数
}

type EmailConfig struct {
	SenderEmail    string
	SenderPassword string
}

func LoadConfig(filePath string) (*Config, error) {
	var config *Config

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
