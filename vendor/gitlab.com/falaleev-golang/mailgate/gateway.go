package mailgate

import (
	"context"
	"fmt"
)

// TemplateRenderer is a basic template renderer interface for subject, html and text body
type TemplateRenderer interface {
	Subject(name string, data map[string]interface{}) (string, error)
	TextBody(name string, data map[string]interface{}) (string, error)
	HtmlBody(name string, data map[string]interface{}) (string, error)
}

// Sender is a sender interface for different senders (mailgun, smtp, etc...)
type Sender interface {
	Send(ctx context.Context, tpl Template) error
}

type MailMessage struct {
	Template string                 `json:"template"`
	Email    string                 `json:"email"`
	Data     map[string]interface{} `json:"data"`
}

// Gateway default way to send email in async way
type Gateway interface {
	Send(ctx context.Context, to interface{}, templateName string, data map[string]interface{}) error
	SendMessage(ctx context.Context, msg *MailMessage) error
}

// MailGateway default way to send email in async way
type gateway struct {
	sender   Sender
	renderer TemplateRenderer
	defaults map[string]interface{}
}

// NewGateway return new gateway instance
func NewGateway(sender Sender, renderer TemplateRenderer, defaults map[string]interface{}) Gateway {
	return &gateway{
		sender:   sender,
		renderer: renderer,
		defaults: defaults,
	}
}

// Send should delivery email to customer or send it to queue
func (g *gateway) render(templateName string, data map[string]interface{}) (subj string, html string, txt string, err error) {
	result := make(map[string]interface{})
	for k, v := range g.defaults {
		result[k] = v
	}
	for k, v := range data {
		result[k] = v
	}

	subj, err = g.renderer.Subject(templateName, result)
	if err != nil {
		return
	}

	html, err = g.renderer.HtmlBody(templateName, result)
	if err != nil {
		return
	}

	txt, err = g.renderer.TextBody(templateName, result)
	return
}

func (g *gateway) getRecipients(to interface{}) ([]string, error) {
	var newTo []string

	switch to.(type) {
	case string:
		newTo = append(newTo, to.(string))
		break

	case []string:
		recipients, ok := to.([]string)
		if ok {
			newTo = append(newTo, recipients...)
		} else {
			return nil, fmt.Errorf("`to` is not a []string")
		}
	}

	return newTo, nil
}

func (g *gateway) Send(ctx context.Context, to interface{}, templateName string, data map[string]interface{}) error {
	newTo, err := g.getRecipients(to)
	if err != nil {
		return err
	}

	subject, bodyHTML, bodyText, err := g.render(templateName, data)
	if err != nil {
		return err
	}

	if len(newTo) == 0 {
		return fmt.Errorf("No recipients")
	}

	return g.sender.Send(ctx, &EmailTemplate{
		to:       newTo,
		subject:  subject,
		bodyHTML: bodyHTML,
		bodyTXT:  bodyText,
	})
}

func (g *gateway) SendMessage(ctx context.Context, msg *MailMessage) error {
	return g.Send(ctx, msg.Email, msg.Template, msg.Data)
}
