package smtp

import (
	"context"
	"gitlab.com/falaleev-golang/mailgate"
	"gopkg.in/gomail.v2"
)

var _ mailgate.Sender = (*smtpSender)(nil)

func New(dialer *gomail.Dialer, from string) mailgate.Sender {
	return &smtpSender{
		dialer: dialer,
		from:   from,
	}
}

type smtpSender struct {
	dialer *gomail.Dialer
	from   string
}

func (s *smtpSender) Send(ctx context.Context, tpl mailgate.Template) error {
	msg := gomail.NewMessage()
	msg.SetHeader("From", s.from)
	msg.SetHeader("To", tpl.To()...)
	msg.SetHeader("Subject", tpl.Subject())
	msg.SetBody("text/plain", tpl.BodyTXT())
	msg.SetBody("text/html", tpl.BodyHTML())

	// Now send E-Mail
	return s.dialer.DialAndSend(msg)
}
