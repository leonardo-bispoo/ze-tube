package bot

import (
	"time"

	"ze-tube/internal/domain"

	"go.mau.fi/whatsmeow"
)

type Handler struct {
	WhatsmeownClient *whatsmeow.Client
	Service          domain.Services
	StartTime        time.Time
	TestUser         string
}
