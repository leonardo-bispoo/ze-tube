package domain

import "go.mau.fi/whatsmeow/types/events"

type (
	Bot interface {
		HandleMessages(evt any)
	}

	Services interface {
		HelloMessage(event *events.Message)
	}
)
