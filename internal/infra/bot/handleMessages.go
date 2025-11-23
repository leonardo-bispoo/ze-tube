package bot

import (
	"go.mau.fi/whatsmeow/types/events"
)

func (h *Handler) HandleMessages(evt any) {
	switch event := evt.(type) {
	case *events.Message:
		if event.Info.IsGroup || event.Info.Timestamp.Before(h.StartTime) {
			return
		}

		if event.Info.Sender.User == h.TestUser {
			h.Service.HelloMessage(event)
		}
	}
}
