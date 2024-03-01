package main

import (
	"io"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	tmdb "github.com/cometbft/cometbft-db"
	tmlog "github.com/cometbft/cometbft/libs/log"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cast"

	"github.com/sentinel-official/hub/v12/app"
	hubtypes "github.com/sentinel-official/hub/v12/types"
)

type appCreator struct {
	encCfg app.EncodingConfig
}

func (ac appCreator) NewApp(
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

	var wasmOpts []wasmkeeper.Option
	if cast.ToBool(appOpts.Get("telemetry.enabled")) {
		wasmOpts = append(wasmOpts, wasmkeeper.WithVMCacheMetrics(prometheus.DefaultRegisterer))
	}

	return app.NewApp(
		appOpts, hubtypes.Bech32MainPrefix, db, ac.encCfg, cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)), true, logger,
		cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants)), skipUpgradeHeights, traceWriter, version.Version,
		wasmOpts, baseapp.SetHaltHeight(cast.ToUint64(appOpts.Get(server.FlagHaltHeight))),
		baseapp.SetHaltTime(cast.ToUint64(appOpts.Get(server.FlagHaltTime))),
		baseapp.SetIndexEvents(cast.ToStringSlice(appOpts.Get(server.FlagIndexEvents))),
		baseapp.SetInterBlockCache(cache),
		baseapp.SetMinGasPrices(cast.ToString(appOpts.Get(server.FlagMinGasPrices))),
		baseapp.SetMinRetainBlocks(cast.ToUint64(appOpts.Get(server.FlagMinRetainBlocks))),
		baseapp.SetPruning(pruningOpts),
		baseapp.SetTrace(cast.ToBool(appOpts.Get(server.FlagTrace))),
	)
}

func (ac appCreator) AppExport(
	logger tmlog.Logger,
	db tmdb.DB,
	traceWriter io.Writer,
	height int64,
	forZeroHeight bool,
	jailWhitelist []string,
	appOpts servertypes.AppOptions,
	modulesToExport []string,
) (servertypes.ExportedApp, error) {
	v := app.NewApp(
		appOpts, hubtypes.Bech32MainPrefix, db, ac.encCfg, cast.ToString(appOpts.Get(flags.FlagHome)),
		cast.ToUint(appOpts.Get(server.FlagInvCheckPeriod)), height == -1, logger,
		cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants)), map[int64]bool{}, traceWriter,
		version.Version, nil, nil,
	)

	if height != -1 {
		if err := v.LoadHeight(height); err != nil {
			return servertypes.ExportedApp{}, err
		}
	}

	return v.ExportAppStateAndValidators(forZeroHeight, jailWhitelist, modulesToExport)
}
