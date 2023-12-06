package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/smtp"
	"time"
)

type MailService struct {
	env     *viper.Viper
	timeout time.Duration
}

func NewMailService(env *viper.Viper) *MailService {
	return &MailService{
		env:     env,
		timeout: time.Duration(4) * time.Second,
	}
}

func (m *MailService) SendMailWithSmtp(ctx context.Context, to []string, subject, body string) error {
	ctx, cancel := context.WithTimeout(ctx, m.timeout)
	defer cancel()

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

func (m *MailService) SendMailWithSendGrid(ctx context.Context, to []string, subject, body string) error {
	panic("implement me")
}
