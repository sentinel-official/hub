package main

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tm "github.com/tendermint/tendermint/types"

	"github.com/sentinel-official/hub/app"
	hubCli "github.com/sentinel-official/hub/app/cli"
	_server "github.com/sentinel-official/hub/server"
	hub "github.com/sentinel-official/hub/types"
)

const (
	flagInvCheckPeriod = "inv-check-period"
)

// nolint:gochecknoglobals
var (
	invCheckPeriod uint
)

func main() {
	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(hub.Bech32PrefixAccAddr, hub.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(hub.Bech32PrefixValAddr, hub.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(hub.Bech32PrefixConsAddr, hub.Bech32PrefixConsPub)
	config.Seal()

	ctx := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "sentinel-hubd",
		Short:             "Sentinel Hub Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	rootCmd.AddCommand(hubCli.InitCmd(ctx, cdc))
	rootCmd.AddCommand(hubCli.CollectGenTxsCmd(ctx, cdc))
	rootCmd.AddCommand(hubCli.TestNetFilesCmd(ctx, cdc))
	rootCmd.AddCommand(hubCli.GenTxCmd(ctx, cdc))
	rootCmd.AddCommand(hubCli.AddGenesisAccountCmd(ctx, cdc))
	rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")

	_server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	executor := cli.PrepareBaseCmd(rootCmd, "SENT_HUB", app.DefaultNodeHome)
	if err := executor.Execute(); err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db tmDB.DB, traceStore io.Writer) abci.Application {
	return app.NewHubApp(
		logger, db, traceStore, true, invCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
	)
}

func exportAppStateAndTMValidators(logger log.Logger, db tmDB.DB, traceStore io.Writer, height int64,
	forZeroHeight bool, jailWhiteList []string) (json.RawMessage, []tm.GenesisValidator, error) {

	if height != -1 {
		hub := app.NewHubApp(logger, db, traceStore, false, uint(1))
		if err := hub.LoadHeight(height); err != nil {
			return nil, nil, err
		}

		return hub.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}

	hub := app.NewHubApp(logger, db, traceStore, true, uint(1))
	return hub.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
