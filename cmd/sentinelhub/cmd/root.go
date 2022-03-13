package cmd

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	clientconfig "github.com/cosmos/cosmos-sdk/client/config"
	"github.com/cosmos/cosmos-sdk/client/debug"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/server"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/snapshots"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcli "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	genutilcli "github.com/cosmos/cosmos-sdk/x/genutil/client/cli"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"

	"github.com/sentinel-official/hub"
	hubparams "github.com/sentinel-official/hub/params"
)

func NewRootCmd() (*cobra.Command, hubparams.EncodingConfig) {
	var (
		config    = hub.MakeEncodingConfig()
		clientCtx = client.Context{}.
				WithCodec(config.Marshaler).
				WithInterfaceRegistry(config.InterfaceRegistry).
				WithTxConfig(config.TxConfig).
				WithLegacyAmino(config.Amino).
				WithInput(os.Stdin).
				WithAccountRetriever(authtypes.AccountRetriever{}).
				WithHomeDir(hub.DefaultNodeHome).
				WithViper("")
	)

	rootCmd := &cobra.Command{
		Use:   "sentinelhub",
		Short: "Sentinel",
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			clientCtx, err := client.ReadPersistentCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			clientCtx, err = clientconfig.ReadFromClientConfig(clientCtx)
			if err != nil {
				return err
			}

			if err = client.SetCmdClientContextHandler(clientCtx, cmd); err != nil {
				return err
			}

			customAppConfigTemplate, customAppConfig := initAppConfig()
			return server.InterceptConfigsPreRunHandler(cmd, customAppConfigTemplate, customAppConfig)
		},
	}

	initRootCmd(rootCmd, config)
	return rootCmd, config
}

func initAppConfig() (string, interface{}) {
	type CustomAppConfig struct {
		serverconfig.Config
	}

	srvCfg := serverconfig.DefaultConfig()
	srvCfg.StateSync.SnapshotInterval = 1000
	srvCfg.StateSync.SnapshotKeepRecent = 10

	appConfig := CustomAppConfig{Config: *srvCfg}
	appConfigTemplate := serverconfig.DefaultConfigTemplate

	return appConfigTemplate, appConfig
}

func initRootCmd(rootCmd *cobra.Command, encodingConfig hubparams.EncodingConfig) {
	cfg := sdk.GetConfig()
	cfg.Seal()

	rootCmd.AddCommand(
		genutilcli.InitCmd(hub.ModuleBasics, hub.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, hub.DefaultNodeHome),
		genutilcli.GenTxCmd(hub.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, hub.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(hub.ModuleBasics),
		AddGenesisAccountCmd(hub.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		testnetCmd(hub.ModuleBasics, banktypes.GenesisBalancesIterator{}),
		debug.Cmd(),
		clientconfig.Cmd(),
	)

	ac := appCreator{
		encCfg: encodingConfig,
	}
	server.AddCommands(rootCmd, hub.DefaultNodeHome, ac.newApp, ac.appExport, addModuleInitFlags)

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
		authcli.GetAccountCmd(),
		rpc.ValidatorCommand(),
		rpc.BlockCommand(),
		authcli.QueryTxsByEventsCmd(),
		authcli.QueryTxCmd(),
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
		authcli.GetSignCommand(),
		authcli.GetSignBatchCommand(),
		authcli.GetMultiSignCommand(),
		authcli.GetMultiSignBatchCmd(),
		authcli.GetValidateSignaturesCommand(),
		flags.LineBreak,
		authcli.GetBroadcastCommand(),
		authcli.GetEncodeCommand(),
		authcli.GetDecodeCommand(),
	)

	hub.ModuleBasics.AddTxCommands(cmd)
	cmd.PersistentFlags().String(flags.FlagChainID, "", "The network chain ID")

	return cmd
}

type appCreator struct {
	encCfg hubparams.EncodingConfig
}

func (ac appCreator) newApp(
	logger log.Logger,
	db tmdb.DB,
	traceStore io.Writer,
	options servertypes.AppOptions,
) servertypes.Application {
	var cache sdk.MultiStorePersistentCache
	if cast.ToBool(options.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(options.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(options)
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
		logger, db, traceStore, true, skipUpgradeHeights,
		cast.ToString(options.Get(flags.FlagHome)),
		cast.ToUint(options.Get(server.FlagInvCheckPeriod)),
		ac.encCfg,
		options,
		baseapp.SetPruning(pruningOpts),
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

func (ac appCreator) appExport(
	logger log.Logger,
	db tmdb.DB,
	traceStore io.Writer,
	height int64,
	forZeroHeight bool,
	jailAllowedAddrs []string,
	appOpts servertypes.AppOptions,
) (servertypes.ExportedApp, error) {
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home is not set")
	}

	var loadLatest bool
	if height == -1 {
		loadLatest = true
	}

	app := hub.NewApp(
		logger,
		db,
		traceStore,
		loadLatest,
		map[int64]bool{},
		homePath,
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		ac.encCfg,
		appOpts,
	)

	if height != -1 {
		if err := app.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return app.ExportAppStateAndValidators(forZeroHeight, jailAllowedAddrs)
}
