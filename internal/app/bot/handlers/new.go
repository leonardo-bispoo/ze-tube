package handlers

import (
	"ze-tube/internal/app/bot"

	"go.mau.fi/whatsmeow"
)

type handler struct {
	WhatsmeownClient *whatsmeow.Client
}

func NewHandlers(client *whatsmeow.Client) bot.Handlers {
	return &handler{
		WhatsmeownClient: client,
	}
}
