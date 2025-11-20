package database

import (
	"context"

	_ "github.com/lib/pq"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

func Connect() *store.Device {
	ctx := context.Background()

	container, err := sqlstore.New(ctx, "postgres", "postgres://localhost:5432/zetube?sslmode=disable", nil)
	if err != nil {
		panic(err)
	}

	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		panic(err)
	}

	return deviceStore
}
