package utils

import (
	"gopkg.in/gomail.v2"
	"github.com/irisnet/iris-community/config"
	"log"
)

func RegisterEmail(email string, id string, code string) error {
	mail := config.Config.Mail
	D := gomail.NewDialer(mail.Host, mail.Port, mail.Username, mail.Password)
	m := gomail.NewMessage()
	m.SetHeader("From", mail.Username)
	m.SetHeader("To", email)                         //收件人
	m.SetHeader("Subject", "账号激活-IRIS")              //标题
	m.SetBody("text/html", "<h1>账号激活</h1><p>尊敬的用户您好，欢迎您注册IRIS</p>"+
		"请在24小时内点击下方按钮或复制下方链接进行邮箱验证"+ id+ " "+ code) //内容

	// Send email
	err := D.DialAndSend(m)
	if err != nil {
		log.Println(err)
	}
	return err
}

func ResetEmail(email string, id string, code string) error {
	mail := config.Config.Mail
	D := gomail.NewDialer(mail.Host, mail.Port, mail.Username, mail.Password)
	m := gomail.NewMessage()
	m.SetHeader("From", mail.Username)
	m.SetHeader("To", email)                         //收件人
	m.SetHeader("Subject", "找回密码-IRIS")              //标题
	m.SetBody("text/html", "<h1>账号激活</h1><p>尊敬的用户您好，欢迎您注册IRIS</p>"+
		"请在24小时内点击下方按钮或复制下方链接进行邮箱验证"+ id+ " "+ code) //内容

	// Send email
	err := D.DialAndSend(m)
	if err != nil {
		log.Println(err)
	}
	return err
}
