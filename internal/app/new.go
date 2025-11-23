package app

import (
	"ze-tube/internal/app/bot"
	"ze-tube/internal/domain"

	"go.mau.fi/whatsmeow"
)

func New(client *whatsmeow.Client) domain.Services {
	return &bot.App{
		WhatsmeownClient: client,
	}
}
