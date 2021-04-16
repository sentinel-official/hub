package cmd

import (
	"io"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/snapshots"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"

	"github.com/sentinel-official/hub"
	"github.com/sentinel-official/hub/params"
)

func NewRootCmd() (*cobra.Command, params.EncodingConfig) {
	var (
		config  = hub.MakeEncodingConfig()
		context = client.Context{}.
			WithJSONMarshaler(config.Marshaler).
			WithInterfaceRegistry(config.InterfaceRegistry).
			WithTxConfig(config.TxConfig).
			WithLegacyAmino(config.Amino).
			WithInput(os.Stdin).
			WithAccountRetriever(types.AccountRetriever{}).
			WithBroadcastMode(flags.BroadcastBlock).
			WithHomeDir(hub.DefaultNodeHome)
	)

	cobra.EnableCommandSorting = false
	root := &cobra.Command{
		Use:   "sentinelhub",
		Short: "Sentinel",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := client.SetCmdClientContextHandler(context, cmd); err != nil {
				return err
			}

			return server.InterceptConfigsPreRunHandler(cmd)
		},
	}

	initRootCmd(root, config)
	return root, config
}

func initRootCmd(rootCmd *cobra.Command, encodingConfig params.EncodingConfig) {
	authclient.Codec = encodingConfig.Marshaler
	rootCmd.AddCommand(
		genutilcli.InitCmd(hub.ModuleBasics, hub.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, hub.DefaultNodeHome),
		migrateCmd(),
		genutilcli.GenTxCmd(hub.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, hub.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(hub.ModuleBasics),
		AddGenesisAccountCmd(hub.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		testnetCmd(hub.ModuleBasics, banktypes.GenesisBalancesIterator{}),
		debug.Cmd(),
	)

	server.AddCommands(rootCmd, hub.DefaultNodeHome, appCreatorFunc, appExportFunc, addModuleInitFlags)
	rootCmd.AddCommand(
		rpc.StatusCommand(),
		queryCommand(),
		txCommand(),
		keys.Commands(hub.DefaultNodeHome),
	)
}
func addModuleInitFlags(startCmd *cobra.Command) {
	crisis.AddModuleInitFlags(startCmd)
}

func queryCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "query",
		Aliases:                    []string{"q"},
		Short:                      "Querying subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcmd.QueryTxsByEventsCmd(),
		authcmd.QueryTxCmd(),
	)

	hub.ModuleBasics.AddQueryCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func txCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        "tx",
		Short:                      "Transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		authcmd.GetSignCommand(),
		authcmd.GetSignBatchCommand(),
		authcmd.GetMultiSignCommand(),
		authcmd.GetMultiSignBatchCmd(),
		authcmd.GetValidateSignaturesCommand(),
		authcmd.GetBroadcastCommand(),
		authcmd.GetEncodeCommand(),
		authcmd.GetDecodeCommand(),
	)

	hub.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

func appCreatorFunc(logger log.Logger, db tmdb.DB, tracer io.Writer, options servertypes.AppOptions) servertypes.Application {
	var cache sdk.MultiStorePersistentCache
	if cast.ToBool(options.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, height := range cast.ToIntSlice(options.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(height)] = true
	}

	pruningOptions, err := server.GetPruningOptionsFromFlags(options)
	if err != nil {
		panic(err)
	}

	snapshotDir := filepath.Join(cast.ToString(options.Get(flags.FlagHome)), "data", "snapshots")
	snapshotDB, err := sdk.NewLevelDB("metadata", snapshotDir)
	if err != nil {
		panic(err)
	}
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	if err != nil {
		panic(err)
	}

	return hub.NewApp(
		logger, db, tracer, true, skipUpgradeHeights,
		cast.ToString(options.Get(flags.FlagHome)),
		cast.ToUint(options.Get(server.FlagInvCheckPeriod)),
		hub.MakeEncodingConfig(),
		options,
		baseapp.SetPruning(pruningOptions),
		baseapp.SetMinGasPrices(cast.ToString(options.Get(server.FlagMinGasPrices))),
		baseapp.SetHaltHeight(cast.ToUint64(options.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(options.Get(server.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(options.Get(server.FlagMinRetainBlocks))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(options.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(options.Get(server.FlagIndexEvents))),
		baseapp.SetSnapshotStore(snapshotStore),
		baseapp.SetSnapshotInterval(cast.ToUint64(options.Get(server.FlagStateSyncSnapshotInterval))),
		baseapp.SetSnapshotKeepRecent(cast.ToUint32(options.Get(server.FlagStateSyncSnapshotKeepRecent))),
	)
}

func appExportFunc(logger log.Logger, db tmdb.DB, tracer io.Writer, height int64,
	forZeroHeight bool, jailAllowedAddrs []string, options servertypes.AppOptions) (servertypes.ExportedApp, error) {
	config := hub.MakeEncodingConfig()
	config.Marshaler = codec.NewProtoCodec(config.InterfaceRegistry)

	var app *hub.App
	if height != -1 {
		app = hub.NewApp(logger, db, tracer, false, map[int64]bool{}, "", cast.ToUint(options.Get(server.FlagInvCheckPeriod)), config, options)

		if err := app.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	} else {
		app = hub.NewApp(logger, db, tracer, true, map[int64]bool{}, "", cast.ToUint(options.Get(server.FlagInvCheckPeriod)), config, options)
	}

	return app.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs)
}
