package main

import (
	"io"
	"path/filepath"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/snapshots"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmdb "github.com/tendermint/tm-db"

	"github.com/sentinel-official/hub/app"
)

type AppCreator struct {
	encCfg  app.EncodingConfig
	homeDir string
}

func (ac AppCreator) NewApp(
	logger tmlog.Logger,
	db tmdb.DB,
	traceWriter io.Writer,
	appOpts servertypes.AppOptions,
) servertypes.Application {
	var cache sdk.MultiStorePersistentCache
	if cast.ToBool(appOpts.Get(server.FlagInterBlockCache)) {
		cache = store.NewCommitKVStoreCacheManager()
	}

	skipUpgradeHeights := make(map[int64]bool)
	for _, height := range cast.ToIntSlice(appOpts.Get(server.FlagUnsafeSkipUpgrades)) {
		skipUpgradeHeights[int64(height)] = true
	}

	pruningOpts, err := server.GetPruningOptionsFromFlags(appOpts)
	if err != nil {
		panic(err)
	}

	snapshotDir := filepath.Join(ac.homeDir, "data", "snapshots")

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

	return app.NewApp(
		appOpts, db, ac.encCfg, ac.homeDir, cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)), true,
		logger, cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants)), skipUpgradeHeights, traceWriter,
		version.Version, wasmOpts, app.GetWasmEnabledProposals(app.DefaultWasmProposals),
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

func (ac AppCreator) AppExport(
	logger tmlog.Logger,
	db tmdb.DB,
	traceWriter io.Writer,
	height int64,
	forZeroHeight bool,
	jailWhitelist []string,
	appOpts servertypes.AppOptions,
) (servertypes.ExportedApp, error) {
	v := app.NewApp(
		appOpts, db, ac.encCfg, ac.homeDir, cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)), height == -1,
		logger, cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants)), map[int64]bool{}, traceWriter,
		version.Version, nil, nil,
	)

	if height != -1 {
		if err := v.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return v.ExportAppStateAndValidators(forZeroHeight, jailWhitelist)
}