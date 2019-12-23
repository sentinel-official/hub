package main

import (
	"encoding/json"
	"io"
	
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	genaccountsCli "github.com/cosmos/cosmos-sdk/x/genaccounts/client/cli"
	genutilCli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tm "github.com/tendermint/tendermint/types"
	db "github.com/tendermint/tm-db"
	
	"github.com/sentinel-official/hub/app"
	_server "github.com/sentinel-official/hub/server"
	hub "github.com/sentinel-official/hub/types"
)

const (
	flagInvCheckPeriod = "inv-check-period"
)

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
	
	rootCmd.AddCommand(genutilCli.InitCmd(ctx, cdc, app.ModuleBasics, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilCli.CollectGenTxsCmd(ctx, cdc, genaccounts.AppModuleBasic{}, app.DefaultNodeHome))
	rootCmd.AddCommand(genutilCli.GenTxCmd(ctx, cdc, app.ModuleBasics, staking.AppModuleBasic{},
		genaccounts.AppModuleBasic{}, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(genutilCli.ValidateGenesisCmd(ctx, cdc, app.ModuleBasics))
	rootCmd.AddCommand(genaccountsCli.AddGenesisAccountCmd(ctx, cdc, app.DefaultNodeHome, app.DefaultCLIHome))
	rootCmd.AddCommand(client.NewCompletionCmd(rootCmd, true))
	
	_server.AddCommands(ctx, cdc, rootCmd, newApp, exportAppStateAndTMValidators)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")
	
	executor := cli.PrepareBaseCmd(rootCmd, "SENT_HUB", app.DefaultNodeHome)
	if err := executor.Execute(); err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db db.DB, traceStore io.Writer) abci.Application {
	return app.NewHubApp(
		logger, db, traceStore, true, invCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(uint64(viper.GetInt(server.FlagHaltHeight))),
	)
}

func exportAppStateAndTMValidators(logger log.Logger, db db.DB, traceStore io.Writer, height int64, forZeroHeight bool,
	jailWhiteList []string) (json.RawMessage, []tm.GenesisValidator, error) {
	if height != -1 {
		hubApp := app.NewHubApp(logger, db, traceStore, false, uint(1))
		err := hubApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}
		return hubApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}
	hubApp := app.NewHubApp(logger, db, traceStore, true, uint(1))
	return hubApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
