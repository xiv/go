package main

import (
	"context"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/network"
	"github.com/stellar/go/services/submitter/internal"
	"github.com/stellar/go/support/db"
)

func main() {
	channels := [1]*internal.Channel{
		{
			Seed: os.Getenv("SUBMITTER_CHANNEL_SEED"),
		},
	}
	sqlxDB, err := sqlx.Connect("postgres", "dbname=submitter sslmode=disable")
	if err != nil {
		return
	}
	store := internal.PostgresStore{
		Session: &db.Session{
			DB: sqlxDB,
		},
	}
	ts := internal.TransactionSubmitter{
		Horizon:         horizonclient.DefaultTestNetClient,
		Network:         network.TestNetworkPassphrase,
		Channels:        channels[:],
		Store:           store,
		RootAccountSeed: os.Getenv("SUBMITTER_ROOT_SEED"),
	}
	ts.Start(context.Background())
}
