package bot

import (
	"ze-tube/internal/domain"

	"go.mau.fi/whatsmeow"
)

type service struct {
	WhatsmeownClient *whatsmeow.Client
}

func New(client *whatsmeow.Client) domain.Services {
	return &service{
		WhatsmeownClient: client,
	}
}
