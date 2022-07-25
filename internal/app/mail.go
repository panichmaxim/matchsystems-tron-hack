package app

import (
	"crypto/tls"
	"embed"
	"github.com/rs/zerolog/log"
	"gitlab.com/falaleev-golang/mailgate/sender/mailgun"
	"gitlab.com/falaleev-golang/mailgate/sender/smtp"
	"gopkg.in/gomail.v2"

	"gitlab.com/falaleev-golang/mailgate"
	"gitlab.com/falaleev-golang/mailgate/render/jetrender"
	"gitlab.com/falaleev-golang/mailgate/sender/chain"
	"gitlab.com/rubin-dev/api/internal/cfg"
	"gitlab.com/rubin-dev/api/pkg/tpl"
)

//go:embed templates/*
var t embed.FS

func createRenderer() (mailgate.TemplateRenderer, error) {
	loader, err := tpl.NewPrefixedFilesystem(t, "templates")
	if err != nil {
		log.Err(err).Msg("error")
		return nil, err
	}
	return jetrender.NewRenderer(loader), nil
}

func createMailGateway(cm cfg.MailgunConfig, cs cfg.SmtpConfig, host string) (mailgate.Gateway, error) {
	loader, err := createRenderer()
	if err != nil {
		log.Err(err).Msg("error")
		return nil, err
	}

	return mailgate.NewGateway(
		chain.New(
			// createSmtpSender(cs),
			createMailgunSender(cm),
		),
		loader,
		map[string]interface{}{"siteUrl": host},
	), nil
}

func createSmtpSender(c cfg.SmtpConfig) mailgate.Sender {
	dialer := gomail.NewDialer(c.SmtpHost, c.SmtpPort, c.SmtpUsername, c.SmtpPassword)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return smtp.New(dialer, c.SmtpFrom)
}

func createMailgunSender(c cfg.MailgunConfig) mailgate.Sender {
	return mailgun.New(c.MailgunDomain, c.MailgunPrivateKey, c.MailgunEndpoint, c.MailgunFrom)
}
