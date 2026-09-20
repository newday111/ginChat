package utils

import (
	"fmt"
	"go_jichu/conf"
	emailformat "go_jichu/internal/utils/emailFormat"

	"github.com/wneessen/go-mail"
)

var EmailSender *emailConfig

type emailConfig struct {
	emailConfig *conf.ConfigStruct
}

func InitEmailConfig(cfg *conf.ConfigStruct) {
	EmailSender = &emailConfig{
		cfg,
	}
}

func (e *emailConfig) SendEmail(toEmail string, code string) error {
	client, err := mail.NewClient(
		e.emailConfig.SMTP.Host,
		mail.WithPort(e.emailConfig.SMTP.Port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(e.emailConfig.SMTP.Username),
		mail.WithPassword(e.emailConfig.SMTP.Password),
	)

	if err != nil {
		return fmt.Errorf("create email stmp client failed: %w", err)
	}

	msg := mail.NewMsg()

	// 发件人
	if err := msg.From(e.emailConfig.SMTP.From); err != nil {
		return fmt.Errorf("set from email failed: %w", err)
	}

	// 收件人
	if err := msg.To(toEmail); err != nil {
		return fmt.Errorf("set to email failed: %w", err)
	}

	msg.Subject("风花雪月注册验证码")

	msg.SetBodyString(
		mail.TypeTextHTML,
		emailformat.RegisterCodeFormat(code),
	)

	if err := client.DialAndSend(msg); err != nil {
		return fmt.Errorf("send email failed: %w", err)
	}
	return nil
}
