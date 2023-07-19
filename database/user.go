package database

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"team_todo/global"
	"team_todo/model"

	"gorm.io/gorm"
)

// 加密
func Myencrypt(row string) string {
	data := []byte(row)
	hash := sha256.Sum256(data)
	fmt.Printf("%x\n", hash)
	res := hex.EncodeToString(hash[:])
	return res
}

// 由service调用的注册函数
func Register(userinfo model.User) error {
	var finduser model.User
	if err := global.GVA_DB.Table("user").
		Where("email = ?", userinfo.Email).
		First(&finduser).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// 用户不存在，可以创建新用户
			userinfo.Password = Myencrypt(userinfo.Password) //加密密码
			if err := global.GVA_DB.Table("user").Create(&userinfo).Error; err != nil {
				return err
			}
			return nil
		}
		// 查询过程中发生错误
		return err
	}
	// 用户已存在，返回错误信息
	return errors.New("EMAIL ALREADY REGISTERD")
}

// 判断指定表tableName的字段column是否存在值value
func Exists(db *gorm.DB, tableName string, column string, value interface{}) bool {

	// 使用map进行动态查询
	condition := map[string]interface{}{
		column: value,
	}

	var count int64
	db.Table(tableName).Where(condition).Count(&count)
	fmt.Println("exist number: ", count)
	return count > 0 //bool count == 0 false-->不存在；count > 0 true-->已存在
}

// 由service调用的登录函数
func Login(userinfo model.User) error {
	//加密用户输入的密码判断是否与数据库中匹配
	var being_logged model.User
	global.GVA_DB.Table("user").Where("email = ?", userinfo.Email).First(&being_logged)
	switch {
	case Myencrypt(userinfo.Password) == being_logged.Password:
		fmt.Println("log in successfully")
		return nil
	case being_logged.Email != "":
		fmt.Println("The email and the password is not matched.")
		//登录失败
		return fmt.Errorf("密码与邮箱不匹配")
	default:
		fmt.Println("The email number is not registered.Please register first.")
		//未注册
		return fmt.Errorf("未注册")
	}
}

// 由service调用的用户信息修改函数
func Modify(userinfo model.User) error {
	var userReq model.User
	userReq.Nickname = userinfo.Nickname
	userReq.Avatar = userinfo.Avatar

	err := global.GVA_DB.Model(&model.User{}).Where("email = ?", userinfo.Email).Updates(&userReq).Error
	if err != nil {
		return err
	}
	return nil
}

// 由service调用的获取用户资料函数
func GetProfile(UserId string) (model.User, error) {
	var user model.User
	err := global.GVA_DB.Table("user").Where("id = ?", UserId).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetId(email string) (string, error) {
	var user model.User
	err := global.GVA_DB.Table("user").Where("email = ?", email).First(&user).Error
	if err != nil {
		return "", err
	}
	return user.Id, nil
}

func UpdateAvatar(id, avatar string) error {
	var user model.User
	if err := global.GVA_DB.Table("user").Where("id = ?", id).First(&user).Error; err != nil {
		return err
	}

	// 修改 avatar 值为指定的 url
	user.Avatar = avatar

	// 使用 Update 方法更新记录
	if err := global.GVA_DB.Table("user").Where("id = ?", id).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}
