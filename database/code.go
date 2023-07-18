package database

import (
	"team_todo/global"
	"team_todo/model"
	"time"

	"gorm.io/gorm"
)

// 把验证码和email存入数据库
func StoreVerCode(db *gorm.DB, code string, expiration time.Time, reqemail string) error {
	verificationCode := model.VerCode{
		Code:       code,
		Expiration: expiration,
		Email:      reqemail,
	}

	result := db.Create(&verificationCode)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 把email对应的验证码从数据库中取出
func GetVerCode(reqemail string) (string, time.Time, error) {
	var being_used model.VerCode
	global.GVA_DB.Table("vercodes").Where("email = ?", reqemail).First(&being_used)

	err := global.GVA_DB.Table("vercodes").Where("email = ?", reqemail).First(&being_used).Error

	if err != nil {
		return "", time.Time{}, err
	}
	return being_used.Code, being_used.Expiration, err

}
