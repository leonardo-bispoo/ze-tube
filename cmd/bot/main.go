package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"ze-tube/internal/app/bot/handlers"
	"ze-tube/internal/infra/database"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
)

func main() {
	deviceStore := database.Connect()

	client := whatsmeow.NewClient(deviceStore, nil)

	handlers := handlers.NewHandlers(client)

	client.AddEventHandler(handlers.HandleMessages)

	if client.Store.ID == nil {
		if err := client.Connect(); err != nil {
			panic(err)
		}

		qrChan, _ := client.GetQRChannel(context.Background())

		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		if err := client.Connect(); err != nil {
			panic(err)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}
