package main

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tm "github.com/tendermint/tendermint/types"
	db "github.com/tendermint/tm-db"

	"github.com/sentinel-official/hub"
	"github.com/sentinel-official/hub/types"
)

const (
	flagInvarCheckPeriod = "invar-check-period"
)

var (
	invarCheckPeriod uint
)

func main() {
	types.GetConfig().Seal()
	cobra.EnableCommandSorting = false

	var (
		cdc = hub.MakeCodec()
		ctx = server.NewDefaultContext()
		cmd = &cobra.Command{
			Use:               "sentinelhubd",
			Short:             "Sentinel Hub Daemon (server)",
			PersistentPreRunE: server.PersistentPreRunEFn(ctx),
		}
	)

	cmd.AddCommand(genutilcli.InitCmd(ctx, cdc, hub.ModuleBasics, hub.DefaultNodeHome))
	cmd.AddCommand(genutilcli.CollectGenTxsCmd(ctx, cdc, auth.GenesisAccountIterator{}, hub.DefaultNodeHome))
	cmd.AddCommand(genutilcli.MigrateGenesisCmd(ctx, cdc))
	cmd.AddCommand(genutilcli.GenTxCmd(ctx, cdc, hub.ModuleBasics, staking.AppModuleBasic{},
		auth.GenesisAccountIterator{}, hub.DefaultNodeHome, hub.DefaultCLIHome))
	cmd.AddCommand(genutilcli.ValidateGenesisCmd(ctx, cdc, hub.ModuleBasics))
	cmd.AddCommand(AddGenesisAccountCmd(ctx, cdc, hub.DefaultNodeHome, hub.DefaultCLIHome))
	cmd.AddCommand(flags.NewCompletionCmd(cmd, true))

	server.AddCommands(ctx, cdc, cmd, newApp, exportAppStateAndValidators)
	cmd.PersistentFlags().UintVar(&invarCheckPeriod, flagInvarCheckPeriod,
		0, "Assert registered invariants every N blocks")

	executor := cli.PrepareBaseCmd(cmd, "SENTINELHUB", hub.DefaultNodeHome)
	if err := executor.Execute(); err != nil {
		panic(err)
	}
}

func newApp(logger log.Logger, db db.DB, tracer io.Writer) abci.Application {
	var cache sdk.MultiStorePersistentCache
	if viper.GetBool(server.FlagInterBlockCache) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	var skipUpgradeHeights = make(map[int64]bool)
	for _, height := range viper.GetIntSlice(server.FlagUnsafeSkipUpgrades) {
		skipUpgradeHeights[int64(height)] = true
	}

	return hub.NewApp(
		logger, db, tracer, true, skipUpgradeHeights, invarCheckPeriod,
		baseapp.SetPruning(storetypes.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
		baseapp.SetHaltHeight(viper.GetUint64(server.FlagHaltHeight)),
		baseapp.SetHaltTime(viper.GetUint64(server.FlagHaltTime)),
		baseapp.SetInterBlockCache(cache),
	)
}

func exportAppStateAndValidators(logger log.Logger, db db.DB, tracer io.Writer, height int64, zeroHeight bool,
	jailWhitelist []string) (json.RawMessage, []tm.GenesisValidator, error) {
	if height != -1 {
		app := hub.NewApp(logger, db, tracer, false, map[int64]bool{}, uint(1))
		if err := app.LoadHeight(height); err != nil {
			return nil, nil, err
		}

		return app.ExportAppStateAndValidators(zeroHeight, jailWhitelist)
	}

	app := hub.NewApp(logger, db, tracer, true, map[int64]bool{}, uint(1))
	return app.ExportAppStateAndValidators(zeroHeight, jailWhitelist)
}
