package mail

import (
	"bytes"
	"html/template"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

var (
	tmpl = template.New("email")
	// SMTP 服务器
	SMTP string
	//Port 端口
	Port int
	// Sender 发件人
	Sender string
	// Password 密码
	Password string
)

// SendEmail 发送邮件
func SendEmail(toPersons []string, subject string, mailTemplate string, data map[string]interface{}) error {
	log.Info("will send mail...")

	m := gomail.NewMessage()
	m.SetHeader("From", "notice@uke.email")
	m.SetHeader("To", toPersons...)
	m.SetHeader("Subject", subject)

	d := gomail.NewDialer(SMTP, Port, Sender, Password)
	tmpl, err := tmpl.Parse(mailTemplate)
	if err != nil {
		return err
	}

	var contents bytes.Buffer
	if err = tmpl.Execute(&contents, data); err != nil {
		return err
	}

	m.SetBody("text/html", contents.String())

	return d.DialAndSend(m)
}
