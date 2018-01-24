package utils

import (
	"gopkg.in/gomail.v2"
	"github.com/irisnet/iris-community/config"
)

func RegisterEmail(to string, name string, code string) {
	mail := config.Config.Mail
	D := gomail.NewDialer(mail.Host, mail.Port, mail.Username, mail.Password)
	m := gomail.NewMessage()
	m.SetHeader("From", mail.Username)
	m.SetHeader("To", to)                                                   //收件人
	m.SetHeader("Subject", "Welcome")                                       //标题
	m.SetBody("text/html", "<h1>Hello,"+name+"!</h1>Welcome to Iris!"+code) //内容

	// Send the email to Bob, Cora and Dan.
	if err := D.DialAndSend(m); err != nil {
		panic(err)
	}
}
