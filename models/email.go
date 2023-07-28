package models

import (
	"fmt"

	"github.com/go-mail/mail/v2"
)

const (
	DefaultSender = "support@aiggato.com"
)

type Email struct {
	From      string
	To        string
	Subject   string
	Plaintext string
	Html      string
}

type EmailService struct {
	DefaultSender string
	dialer        *mail.Dialer
}

type SMTPConfig struct {
	Host     string
	Port     int
	Username string
	Password string
}

func NewEmailService(config SMTPConfig) *EmailService {
	es := EmailService{
		dialer: mail.NewDialer(config.Host, config.Port, config.Username, config.Password),
	}
	return &es

}

func (es *EmailService) SendEmail(e Email) error {
	msg := mail.NewMessage()
	es.setFrom(msg, e)
	msg.SetHeader("To", e.To)
	msg.SetHeader("Subject", e.Subject)
	switch {
	case e.Plaintext != "" && e.Html != "":
		msg.SetBody("text/plain", e.Plaintext)
		msg.AddAlternative("text/html", e.Html)
	case e.Plaintext != "":
		msg.SetBody("text/plain", e.Plaintext)
	case e.Html != "":
		msg.AddAlternative("text/html", e.Html)
	}

	err := es.dialer.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("send: %w", err)
	}
	return nil
}

func (es *EmailService) ForgotPassword(to, resetURL string) error {
	email := Email{
		Subject:   "Reset your password",
		To:        to,
		Plaintext: "Click here to reset your password: " + resetURL,
		Html:      `<p>To reset your password, please follow the link: <a href="` + resetURL + `">` + resetURL + `</a></p>`,
	}
	err := es.SendEmail(email)
	if err != nil {
		return fmt.Errorf("forgot password email: %w", err)
	}
	return nil
}

func (es *EmailService) setFrom(msg *mail.Message, email Email) {
	var from string
	switch {
	case email.From != "":
		from = email.From
	case es.DefaultSender != "":
		from = es.DefaultSender
	default:
		from = DefaultSender
	}
	msg.SetHeader("From", from)
}
