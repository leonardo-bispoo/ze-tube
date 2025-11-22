package bot

import (
	"ze-tube/internal/app/bot"
	"ze-tube/internal/domain"

	"go.mau.fi/whatsmeow"
)

type handler struct {
	WhatsmeownClient *whatsmeow.Client
	Service          domain.Services
}

func New(client *whatsmeow.Client) domain.Bot {
	return &handler{
		Service: bot.New(client),
	}
}
