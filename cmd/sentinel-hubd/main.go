package main

import (
	"encoding/json"
	"io"
	"os"

	"github.com/cosmos/cosmos-sdk/baseapp"

	"github.com/cosmos/cosmos-sdk/server"
	"github.com/ironman0x7b2/sentinel-hub/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmTypes "github.com/tendermint/tendermint/types"
)

func main() {
	cdc := app.MakeCodec()
	ctx := server.NewDefaultContext()

	rootCmd := &cobra.Command{
		Use:               "sentinel-hubd",
		Short:             "Sentinel Hub Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	server.AddCommands(ctx, cdc, rootCmd, server.DefaultAppInit,
		server.ConstructAppCreator(newApp, "sentinel-hub"),
		server.ConstructAppExporter(exportAppStateAndTMValidators, "sentinel-hub"))

	rootDir := os.ExpandEnv("$HOME/.sentinel-hubd")
	executor := cli.PrepareBaseCmd(rootCmd, "BC", rootDir)

	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db tmDB.DB, storeTracer io.Writer) abciTypes.Application {
	return app.NewSentinelHub(logger, db, baseapp.SetPruning(viper.GetString("pruning")))
}

func exportAppStateAndTMValidators(logger log.Logger, db tmDB.DB, storeTracer io.Writer) (json.RawMessage, []tmTypes.GenesisValidator, error) {
	bapp := app.NewSentinelHub(logger, db)

	return bapp.ExportAppStateAndValidators()
}
