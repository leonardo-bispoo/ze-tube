package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ze-tube/internal/config"
	"ze-tube/internal/config/database/postgres"
	"ze-tube/internal/config/database/redis"

	infra "ze-tube/internal/infra"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal/v3"
	"go.mau.fi/whatsmeow"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func main() {
	config := config.LoadEnv()

	ctx := context.Background()

	deviceStore := postgres.ConectDevice(ctx, config.PostgresURL)

	redisStore := redis.Connect(ctx, config.RedisURL)

	defer redisStore.Close()

	clientLog := waLog.Stdout("Client", "DEBUG", true)

	client := whatsmeow.NewClient(deviceStore, clientLog)

	now := time.Now()

	botHandlers := infra.NewHandler(client, now)

	client.AddEventHandler(botHandlers.HandleMessages)

	if client.Store.ID == nil {
		qrChan, _ := client.GetQRChannel(context.Background())

		if err := client.Connect(); err != nil {
			panic(err)
		}

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
