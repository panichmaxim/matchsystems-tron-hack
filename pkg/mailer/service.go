package mailer

import (
	"context"
	"fmt"
	"time"

	"gitlab.com/falaleev-golang/mailgate"
)

type UserNotify interface {
	RegistrationConfirm(ctx context.Context, email string) error
	RestoreConfirm(ctx context.Context, email string) error
	Registration(ctx context.Context, email, token string) error
	Restore(ctx context.Context, email, token string) error
}

type AccessTokenNotify interface {
	AccessRequest(ctx context.Context, to, name, email string) error
}

func NewNotify(gate mailgate.Gateway) Notify {
	return &notify{gate: gate}
}

type Notify interface {
	UserNotify
	AccessTokenNotify
}

type notify struct {
	gate mailgate.Gateway
}

func (n *notify) AccessRequest(ctx context.Context, to, name, email string) error {
	return n.gate.SendMessage(ctx, &mailgate.MailMessage{
		Template: "notification",
		Email:    to,
		Data: map[string]interface{}{
			"date":  time.Now(),
			"title": "Access request",
			"text": fmt.Sprintf(
				"user %s (%s) want to access to paid functionality",
				name,
				email,
			),
		},
	})
}

func (n *notify) RegistrationConfirm(ctx context.Context, email string) error {
	return n.gate.SendMessage(ctx, &mailgate.MailMessage{
		Template: "notification",
		Email:    email,
		Data: map[string]interface{}{
			"date":  time.Now(),
			"title": "Подтверждение регистрации",
			"text":  "Благодарим вас за регистрацию",
		},
	})
}

func (n *notify) RestoreConfirm(ctx context.Context, email string) error {
	return n.gate.SendMessage(ctx, &mailgate.MailMessage{
		Template: "notification",
		Email:    email,
		Data: map[string]interface{}{
			"date":  time.Now(),
			"title": "Изменение пароля",
			"text":  "Пароль успешно изменен",
		},
	})
}

func (n *notify) Registration(ctx context.Context, email, token string) error {
	return n.gate.SendMessage(ctx, &mailgate.MailMessage{
		Template: "action",
		Email:    email,
		Data: map[string]interface{}{
			"date":       time.Now(),
			"title":      "Регистрация",
			"text":       "Для подтверждения регистрации перейдите по ссылке из письма",
			"actionText": "Продолжить регистрацию",
			"actionUrl":  fmt.Sprintf("/auth/registration/%s", token),
		},
	})
}

func (n *notify) Restore(ctx context.Context, email, token string) error {
	return n.gate.SendMessage(ctx, &mailgate.MailMessage{
		Template: "action",
		Email:    email,
		Data: map[string]interface{}{
			"date":       time.Now(),
			"title":      "Восстановление пароля",
			"text":       "Для восстановления пароля перейдите по ссылке из письма",
			"actionText": "Восстановить пароль",
			"actionUrl":  fmt.Sprintf("/auth/restore/%s", token),
		},
	})
}
