package mailgun

import (
	"context"

	api "github.com/mailgun/mailgun-go/v3"
	"gitlab.com/falaleev-golang/mailgate"
)

var _ mailgate.Sender = (*mailgunSender)(nil)

// New new mailgun sender
func New(domain, privateKey, endpoint, from string) mailgate.Sender {
	m := api.NewMailgun(domain, privateKey)
	m.SetAPIBase(endpoint)

	return &mailgunSender{m, from}
}

type mailgunSender struct {
	mg   *api.MailgunImpl
	from string
}

// Send email message via mailgun
func (m *mailgunSender) Send(ctx context.Context, tpl mailgate.Template) error {
	msg := m.mg.NewMessage(m.from, tpl.Subject(), tpl.BodyTXT(), tpl.To()...)
	msg.SetHtml(tpl.BodyHTML())

	_, _, err := m.mg.Send(ctx, msg)

	return err
}
