package postgres

import (
	"context"

	_ "github.com/lib/pq"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/store/sqlstore"
)

func ConectDevice(ctx context.Context, url string) *store.Device {
	container, err := sqlstore.New(ctx, "postgres", url, nil)
	if err != nil {
		panic(err)
	}

	deviceStore, err := container.GetFirstDevice(ctx)
	if err != nil {
		panic(err)
	}

	return deviceStore
}
