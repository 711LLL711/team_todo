package model

import (
	"time"

	//"gorm.io/gorm"
)

//定义验证码结构体

type VerCode struct {
	Code       string
	Expiration time.Time
	Email      string
}
