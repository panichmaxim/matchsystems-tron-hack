package chain

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"gitlab.com/falaleev-golang/mailgate"
)

var _ mailgate.Sender = (*chainSender)(nil)

// New new mailgun sender
func New(senders ...mailgate.Sender) mailgate.Sender {
	return &chainSender{senders}
}

type chainSender struct {
	senders []mailgate.Sender
}

// Send email message via mailgun
func (m *chainSender) Send(ctx context.Context, tpl mailgate.Template) error {
	for _, s := range m.senders {
		err := s.Send(ctx, tpl)
		if err == nil {
			return nil
		}
		if err != nil {
			log.Err(err).Msg("error while send email, try next sender")
		}
	}

	return fmt.Errorf("all senders in failed state")
}
