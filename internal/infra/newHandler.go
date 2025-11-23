package bot

import (
	"os"
	"time"

	"ze-tube/internal/app"
	"ze-tube/internal/domain"
	"ze-tube/internal/infra/bot"

	"go.mau.fi/whatsmeow"
)

func NewHandler(client *whatsmeow.Client, startTime time.Time) domain.Bot {
	testUser := os.Getenv("TEST_USER_TO_SEND_MESSAGE")

	return &bot.Handler{
		Service:   app.New(client),
		StartTime: startTime,
		TestUser:  testUser,
	}
}
