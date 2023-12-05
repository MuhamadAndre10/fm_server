package service

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/smtp"
)

type MailService struct {
	env *viper.Viper
}

func NewMailService(env *viper.Viper) *MailService {
	return &MailService{
		env: env,
	}
}

func (m *MailService) SendMailWithSmtp(to []string, subject, body string) error {
	user := "andrepriyanto95@gmail.com"
	pass := m.env.GetString("MAIL_PASS")
	host := m.env.GetString("MAIL_HOST")
	port := m.env.GetString("MAIL_PORT")

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	sub := "Subject: " + subject + "!\n"
	msg := []byte(sub + mime + "\n" + body)
	addr := host + ":" + port

	auth := smtp.PlainAuth("", user, pass, host)

	err := smtp.SendMail(addr, auth, user, to, msg)
	if err != nil {
		return errors.WithMessage(err, "failed to send email")
	}

	return nil
}

func (m *MailService) SendMailWithSendGrid(to []string, subject, body string) error {
	panic("implement me")
}
