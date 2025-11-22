package bot

import "go.mau.fi/whatsmeow/types/events"

func (h *handler) HandleMessages(evt any) {
	switch event := evt.(type) {
	case *events.Message:
		h.Service.HelloMessage(event)
	}
}
