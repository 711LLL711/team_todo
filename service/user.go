package service

import (
	"team_todo/config"
	"team_todo/database"
	"team_todo/global"
	"team_todo/model"

	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"strings"
	"team_todo/util"
	"time"
)

// 用户注册
func Register(userinfo model.User) error {
	// 在这里调用数据库包中的函数来进行用户注册逻辑处理

	err := database.Register(userinfo)
	return err
}

// 用户登录
func Login(userinfo model.User, req *http.Request, resp http.ResponseWriter) error {
	// 在这里调用数据库包中的函数来进行用户登录逻辑处理
	session.Email = userinfo.Email
	session.Id, err = database.GetId(userinfo.Email)
	if err != nil {
		return err
	}
	err = util.SetSession(resp, req, session.Id, session.Email)
	if err != nil {
		return err
	}
	return nil
}

// 用户信息修改
func Modify(userinfo model.User) error {
	// 在这里调用数据库包中的函数来进行用户信息修改逻辑处理
	err := database.Modify(userinfo)
	return err

// 读取文件配置，获得
func SenderEmail() (*config.EmailConfig, error) {
	configFilePath := "../config/config.json"
	config, err := config.LoadConfig(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	return &config.Email, nil
}

// 发送验证邮件
// func SendVerCodeByEmail(email, code string) error {
// 	senderemail, err1 := senderemail()
// 	if err1 != nil {
// 		return err1
// 	}
// 	// 创建邮件对象
// 	m := gomail.NewMessage()
// 	m.SetHeader("From", senderEmail.SenderEmail)          // 发件人
// 	m.SetHeader("To", email)                              // 收件人
// 	m.SetHeader("Subject", "Verification Code for Login") // 邮件主题
// 	m.SetBody("text/plain", "Your verification code for Team_todo is "+code)

// 	// 配置SMTP服务器信息
// 	d := gomail.NewDialer(senderEmail.SMTP_server, 587, senderEmail.SenderEmail, senderEmail.SenderPassword)

// 	// 发送邮件
// 	err := d.DialAndSend(m)
// 	if err != nil {
// 		return err
// 	}

//		return nil
//	}

// 邮件发送格式化函数
func SendMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth(password, user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

// 邮件发送执行函数
func PerformEmailSending(reqemail, code string) error {
	// s := GenRanNum(100000, 999999)
	// fmt.Println(s)

	senderemail, err1 := SenderEmail()
	if err1 != nil {
		return err1
	}
	user := senderemail.SenderEmail
	password := senderemail.SenderPassword
	host := "smtp.qq.com:587"
	to := reqemail
	subject := "Verification Code from Team_todo"
	//     body := `
	//  <html>
	//  <body>
	//  <h3>
	//  "Test send email by golang"
	//  </h3>
	//  </body>
	//  </html>
	//  `
	body := "Your verification code for Team_todo is " + code + ". Valid for 5 min"
	fmt.Println("send email")
	err := SendMail(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Println("send mail error!")
		fmt.Println(err)
		return err
	} else {
		fmt.Println("send mail success!")
	}
	return nil
}

// 生成六位验证码，和邮箱一起存入数据库
func GenVerCode(reqemail string) string {
	//生成六位随机数
	code := GenRanNum(100000, 999999)

	// 设置验证码的有效期为5分钟
	var verificationCode model.VerCode
	verificationCode.Expiration = time.Now().Add(5 * time.Minute)

	// 将验证码和有效期存储在数据库或缓存中，以便后续验证时使用
	database.StoreVerCode(global.GVA_DB, code, verificationCode.Expiration, reqemail)
	return code
}

// 生成六位随机数（已验）
func GenRanNum(min, max int) string {
	// 设置随机数种子
	seed := time.Now().UnixNano()
	rand := rand.New(rand.NewSource(seed))
	// 生成指定范围内的随机数
	number := rand.Intn(max-min+1) + min

	// 格式化为六位数的字符串
	return fmt.Sprintf("%06d", number)
}

// 检查验证码是否正确
func CheckVercode(code string, reqemail string) bool {
	// 从数据库或缓存中获取验证码和有效期
	var being_used model.VerCode
	var err error
	being_used.Code, being_used.Expiration, err = database.GetVerCode(reqemail)
	if err != nil {
		return false
	}
	// 检查当前时间是否在有效期内
	currentTime := time.Now()
	if currentTime.Before(being_used.Expiration) {
		// 执行验证码的验证逻辑
		if code == being_used.Code {
			return true
		}
	}

	return false
}

//获取用户资料

func GetProfile(UserId string) (model.User, error) {
	user, err := database.GetProfile(UserId)
	return user, err
}
