package handlers

import (
	"context"

	"ze-tube/internal/app/bot/services"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func (h *handler) HandleMessages(evt any) {
	switch v := evt.(type) {
	case *events.Message:
		message := v.Message.GetConversation()

		ctx := context.Background()

		h.WhatsmeownClient.SendMessage(ctx, v.Info.Chat, &waE2E.Message{
			Conversation: proto.String(`Olá! Eu sou o Zé Tube — seu assistente para baixar vídeos e áudios do YouTube 🎧

Envie um link do YouTube e eu baixo pra você rapidinho:
- 🎬 Vídeo em qualidade média
- ♬ Áudio em MP3

Exemplo:
https://youtu.be/dQw4w9WgXcQ`),
		})

		if !services.IsValidURL(message) {
			h.WhatsmeownClient.SendMessage(ctx, v.Info.Chat, &waE2E.Message{
				Conversation: proto.String("Por favor, envie um link de video de youtube válido."),
			})
			return
		}
	}
}
