package main

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
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
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	tmcli "github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"

	"github.com/sentinel-official/hub"
	hubparams "github.com/sentinel-official/hub/params"
	hubtypes "github.com/sentinel-official/hub/types"
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
				WithViper("SENTINEL")
	)

	rootCmd := &cobra.Command{
		Use:   "sentinelhub",
		Short: "Sentinel Hub application",
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
	srvCfg.BaseConfig.MinGasPrices = "0.1udvpn"
	srvCfg.StateSync.SnapshotInterval = 1000

	appConfig := CustomAppConfig{Config: *srvCfg}
	appConfigTemplate := serverconfig.DefaultConfigTemplate

	return appConfigTemplate, appConfig
}

func initRootCmd(rootCmd *cobra.Command, encodingConfig hubparams.EncodingConfig) {
	cfg := hubtypes.GetConfig()
	cfg.Seal()

	rootCmd.AddCommand(
		genutilcli.InitCmd(hub.ModuleBasics, hub.DefaultNodeHome),
		genutilcli.CollectGenTxsCmd(banktypes.GenesisBalancesIterator{}, hub.DefaultNodeHome),
		genutilcli.GenTxCmd(hub.ModuleBasics, encodingConfig.TxConfig, banktypes.GenesisBalancesIterator{}, hub.DefaultNodeHome),
		genutilcli.ValidateGenesisCmd(hub.ModuleBasics),
		AddGenesisAccountCmd(hub.DefaultNodeHome),
		AddGenesisWasmMsgCmd(hub.DefaultNodeHome),
		tmcli.NewCompletionCmd(rootCmd, true),
		debug.Cmd(),
		clientconfig.Cmd(),
	)

	ac := appCreator{encCfg: encodingConfig}
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
	wasm.AddModuleInitFlags(startCmd)
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
	tracer io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	var cache sdk.MultiStorePersistentCache
	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, h := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(h)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	snapshotDir := filepath.Join(cast.ToString(appOpts.Get(flags.FlagHome)), "data", "snapshots")
	snapshotDB, err := sdk.NewLevelDB("metadata", snapshotDir)
	if err != nil {
		panic(err)
	}
	snapshotStore, err := snapshots.NewStore(snapshotDB, snapshotDir)
	if err != nil {
		panic(err)
	}

	var wasmOpts []wasmkeeper.Option
	if cast.ToBool(appOpts.Get("telemetry.enabled")) {
		wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
	}

	return hub.NewApp(
		logger, db, tracer, true, skipUpgradeHeights,
		cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		ac.encCfg,
		hub.GetWasmEnabledProposals(),
		appOpts,
		wasmOpts,
		baseapp.SetPruning(pruningOpts),
		baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
		baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
		baseapp.SetSnapshotStore(snapshotStore),
		baseapp.SetSnapshotInterval(cast.ToUint64(appOpts.Get(server.FlagStateSyncSnapshotInterval))),
		baseapp.SetSnapshotKeepRecent(cast.ToUint32(appOpts.Get(server.FlagStateSyncSnapshotKeepRecent))),
	)
}

func (ac appCreator) appExport(
	logger log.Logger,
	db tmdb.DB,
	tracer io.Writer,
	height int64,
	forZeroHeight bool,
	jailWhitelist []string,
	appOpts servertypes.AppOptions,
) (servertypes.ExportedApp, error) {
	homePath, ok := appOpts.Get(flags.FlagHome).(string)
	if !ok || homePath == "" {
		return servertypes.ExportedApp{}, errors.New("application home is not set")
	}

	app := hub.NewApp(
		logger,
		db,
		tracer,
		height == -1,
		map[int64]bool{},
		homePath,
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)),
		ac.encCfg,
		hub.GetWasmEnabledProposals(),
		appOpts,
		[]wasmkeeper.Option{},
	)

	if height != -1 {
		if err := app.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return app.ExportAppStateAndValidators(forZeroHeight, jailWhitelist)
}
