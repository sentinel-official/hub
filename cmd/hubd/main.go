package main

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	csdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tm "github.com/tendermint/tendermint/types"

	app "github.com/ironman0x7b2/sentinel-sdk/app/hub"
	hubCli "github.com/ironman0x7b2/sentinel-sdk/app/hub/cli"
	_server "github.com/ironman0x7b2/sentinel-sdk/server"
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

	config := csdk.GetConfig()
	config.SetBech32PrefixForAccount(csdk.Bech32PrefixAccAddr, csdk.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(csdk.Bech32PrefixValAddr, csdk.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(csdk.Bech32PrefixConsAddr, csdk.Bech32PrefixConsPub)
	config.Seal()

	ctx := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "hubd",
		Short:             "Hub Daemon (server)",
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

	executor := cli.PrepareBaseCmd(rootCmd, "HUB", app.DefaultNodeHome)
	if err := executor.Execute(); err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db tmDB.DB, traceStore io.Writer) abci.Application {
	return app.NewHub(
		logger, db, traceStore, true, invCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
	)
}

func exportAppStateAndTMValidators(logger log.Logger, db tmDB.DB, traceStore io.Writer, height int64,
	forZeroHeight bool, jailWhiteList []string) (json.RawMessage, []tm.GenesisValidator, error) {

	if height != -1 {
		hub := app.NewHub(logger, db, traceStore, false, uint(1))
		if err := hub.LoadHeight(height); err != nil {
			return nil, nil, err
		}

		return hub.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}

	hub := app.NewHub(logger, db, traceStore, true, uint(1))
	return hub.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
