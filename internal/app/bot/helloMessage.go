package bot

import (
	"context"

	"go.mau.fi/whatsmeow/proto/waE2E"
	"go.mau.fi/whatsmeow/types/events"
	"google.golang.org/protobuf/proto"
)

func (s *service) HelloMessage(event *events.Message) {
	ctx := context.Background()

	s.WhatsmeownClient.SendMessage(ctx, event.Info.Chat, &waE2E.Message{
		Conversation: proto.String(`Olá! Eu sou o Zé Tube — seu assistente para baixar vídeos e áudios do YouTube 🎧

Envie um link do YouTube e eu baixo pra você rapidinho:
- 🎬 Vídeo em qualidade média
- ♬ Áudio em MP3

Exemplo:
https://youtu.be/dQw4w9WgXcQ`),
	})
}
