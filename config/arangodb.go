package config

import (
	"context"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

func NewArangoDBDatabase(URL, user, pass string) driver.Client {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{URL}, //DB Url
	})
	if err != nil {
		panic(err)
	}

	client, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(user, pass),
	})
	if err != nil {
		panic(err)
	}

	return client
}

func NewArangoDBContext(timeout time.Duration) (context.Context, context.CancelFunc) {
	var ctx = context.Background()
	return context.WithTimeout(ctx, timeout)
}
