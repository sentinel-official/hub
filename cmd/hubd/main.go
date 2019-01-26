package main

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmTypes "github.com/tendermint/tendermint/types"

	app "github.com/ironman0x7b2/sentinel-sdk/apps/hub"
	hubCli "github.com/ironman0x7b2/sentinel-sdk/apps/hub/cli"
)

func main() {
	cdc := app.MakeCodec()

	config := csdkTypes.GetConfig()
	config.SetBech32PrefixForAccount(csdkTypes.Bech32PrefixAccAddr, csdkTypes.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(csdkTypes.Bech32PrefixValAddr, csdkTypes.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(csdkTypes.Bech32PrefixConsAddr, csdkTypes.Bech32PrefixConsPub)
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
	rootCmd.AddCommand(hubCli.TestnetFilesCmd(ctx, cdc))
	rootCmd.AddCommand(hubCli.GenTxCmd(ctx, cdc))
	rootCmd.AddCommand(hubCli.AddGenesisAccountCmd(ctx, cdc))
	rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))

	server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)

	executor := cli.PrepareBaseCmd(rootCmd, "SH", app.DefaultNodeHome)
	err := executor.Execute()
	if err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db tmDB.DB, traceStore io.Writer) abciTypes.Application {
	return app.NewHub(
		logger, db, traceStore, true,
		baseapp.SetPruning(store.NewPruningOptions(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
	)
}

func exportAppStateAndTMValidators(logger log.Logger, db tmDB.DB, traceStore io.Writer, height int64,
	forZeroHeight bool) (json.RawMessage, []tmTypes.GenesisValidator, error) {
	if height != -1 {
		hub := app.NewHub(logger, db, traceStore, false)
		err := hub.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return hub.ExportAppStateAndValidators(forZeroHeight)
	}
	hub := app.NewHub(logger, db, traceStore, true)
	return hub.ExportAppStateAndValidators(forZeroHeight)
}
