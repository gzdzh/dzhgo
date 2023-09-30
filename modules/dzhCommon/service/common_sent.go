package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gzdzh/dzhgo/dzhCore"

	"gopkg.in/gomail.v2"
)

type CommonSentService struct {
	*dzhCore.Service
}

// 发送邮件
func (s *CommonSentService) SentEmail(content string, subject string, addressHeader string, config gdb.Record) {

	var ctx context.Context
	// 邮件开启
	if config["remindEmail"].Int() == 1 {
		// 设置 SMTP 服务器的认证信息
		smtp := config["smtp"].String()
		smtpPort := 465
		senderEmail := config["sendEmail"].String()
		senderPassword := config["pass"].String()

		body := content
		// 邮件内容
		toEmail := config["requestEmail"].String()
		toEmails := strings.Split(toEmail, "|")

		m := gomail.NewMessage()
		m.SetHeader("To", toEmails...)
		m.SetHeader("Subject", subject)
		m.SetAddressHeader("From", senderEmail, addressHeader)
		m.SetBody("text/html", body)

		d := gomail.NewDialer(smtp, smtpPort, senderEmail, senderPassword)
		// 发送
		err := d.DialAndSend(m)
		if err != nil {
			panic(err)
		}
		g.Log().Debug(ctx, "邮件发送成功")
		fmt.Println("发送成功")
	}

}

func NewCommonSentService() *CommonSentService {
	return &CommonSentService{
		&dzhCore.Service{},
	}
}
