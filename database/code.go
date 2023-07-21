package database

import (
	"log"
	"team_todo/global"
	"team_todo/model"

	//"time"

	"gorm.io/gorm"
)

// 把验证码和email存入数据库
func StoreVerCode(db *gorm.DB, code string, expiration string, reqemail string) error {
	verificationCode := model.VerCode{
		Code:       code,
		Expiration: expiration,
		Email:      reqemail,
	}

	result := db.Table("vercode").Create(&verificationCode)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 把email对应的验证码从数据库中取出
func GetVerCode(reqemail string) (string, string, error) {
	var being_used model.VerCode
	global.GVA_DB.Table("vercode").Where("email = ?", reqemail).Last(&being_used)
result := global.GVA_DB.Table("vercode").Where("email = ?", reqemail).Order("expiration DESC").First(&being_used)
	log.Println("database code :",being_used)
	

	if result.Error != nil {
		return "", "", result.Error
	}
	return being_used.Code, being_used.Expiration, nil

}
