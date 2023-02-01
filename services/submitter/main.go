package main

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/stellar/go/clients/horizonclient"
	"github.com/stellar/go/network"
	"github.com/stellar/go/services/submitter/internal"
	"github.com/stellar/go/support/db"
)

func main() {
	channels := [1]*internal.Channel{
		{
			Seed: "GBEKPFYEFY76INT7JVTZCTIFB3MS5KM3T3RTE2GKP2I6THVHOFZKJHVH",
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
		RootAccountSeed: "SBSYXCIQPAG5NOWXEIUM3665VIOZK4AUSDRI5ARQRTQZXADD5PBIGJZ3",
	}
	ts.Start(context.Background())
}
