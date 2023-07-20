package reminder

import (
	"fmt"
	"net/smtp"
	"strings"
	"team_todo/config"
	"team_todo/model"
)

//具体的提醒逻辑

// 大概率要debug
func SenderEmail() (*config.EmailConfig, error) {
	configFilePath := "../config/config.json"
	config, err := config.LoadConfig(configFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	return &config.Email, nil
}
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
func Conduct(reminder model.Reminder) error {

	// s := GenRanNum(100000, 999999)
	// fmt.Println(s)

	senderemail, err1 := SenderEmail()
	if err1 != nil {
		return err1
	}
	user := senderemail.SenderEmail
	password := senderemail.SenderPassword
	host := "smtp.qq.com:587"
	to := reminder.Email
	subject := "Reminder from Team_todo"
	//     body := `
	//  <html>
	//  <body>
	//  <h3>
	//  "Test send email by golang"
	//  </h3>
	//  </body>
	//  </html>
	//  `
	body := "Dear " + reminder.Nickname + ",your task " + reminder.TaskId + " is due by " + reminder.Deadline + ". "
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
