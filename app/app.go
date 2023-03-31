package app

import (
	"encoding/json"
	"io"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server/api"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	abcitypes "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmlog "github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	"github.com/sentinel-official/hub/app/ante"
	"github.com/sentinel-official/hub/app/upgrades"
)

const (
	appName = "Sentinel Hub"
)

var (
	_ simapp.App              = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)

type App struct {
	*baseapp.BaseApp
	EncodingConfig
	Keepers
	StoreKeys
	mm *module.Manager
	sm *module.SimulationManager
}

func NewApp(
	appOpts servertypes.AppOptions,
	db tmdb.DB,
	encCfg EncodingConfig,
	homeDir string,
	invCheckPeriod uint,
	loadLatest bool,
	logger tmlog.Logger,
	skipGenesisInvariants bool,
	skipUpgradeHeights map[int64]bool,
	traceWriter io.Writer,
	version string,
	wasmOpts []wasmkeeper.Option,
	wasmProposalTypes []wasmtypes.ProposalType,
	baseAppOpts ...func(*baseapp.BaseApp),
) *App {
	baseApp := baseapp.NewBaseApp(appName, logger, db, encCfg.TxConfig.TxDecoder(), baseAppOpts...)
	baseApp.SetCommitMultiStoreTracer(traceWriter)
	baseApp.SetVersion(version)
	baseApp.SetInterfaceRegistry(encCfg.InterfaceRegistry)

	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic("failed to read the wasm config: " + err.Error())
	}

	var (
		storeKeys = NewStoreKeys()
		keepers   = NewKeepers(
			encCfg.Amino, baseApp, BlockedAccAddrs(), encCfg.Codec, homeDir, invCheckPeriod,
			storeKeys, ModuleAccPerms(), skipUpgradeHeights, wasmConfig, wasmOpts, wasmProposalTypes,
		)
		mm = NewModuleManager(
			encCfg.Codec, baseApp.DeliverTx, encCfg.InterfaceRegistry, keepers, skipGenesisInvariants, encCfg.TxConfig,
		)
		sm = NewSimulationManager(encCfg.Codec, encCfg.InterfaceRegistry, keepers)
	)

	app := &App{
		BaseApp:        baseApp,
		EncodingConfig: encCfg,
		Keepers:        keepers,
		StoreKeys:      storeKeys,
		mm:             mm,
		sm:             sm,
	}

	app.mm.RegisterInvariants(&keepers.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encCfg.Amino)

	configurator := module.NewConfigurator(encCfg.Codec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(configurator)

	app.MountKVStores(app.KVKeys())
	app.MountMemoryStores(app.MemoryKeys())
	app.MountTransientStores(app.TransientKeys())

	app.SetupAnteHandler(wasmConfig)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.SetInitChainer(app.InitChainer)
	app.SetUpgradeHandler(configurator)
	app.SetUpgradeStoreLoader()
	app.RegisterSnapshotExtensions()

	if loadLatest {
		if err = app.LoadLatestVersion(); err != nil {
			tmos.Exit("failed to load the latest version: " + err.Error())
		}

		ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})
		if err = app.WasmKeeper.InitializePinnedCodes(ctx); err != nil {
			tmos.Exit("failed to initialize the pinned codes: " + err.Error())
		}
	}

	return app
}

func (a *App) LegacyAmino() *codec.LegacyAmino {
	return a.Amino
}

func (a *App) BeginBlocker(ctx sdk.Context, req abcitypes.RequestBeginBlock) abcitypes.ResponseBeginBlock {
	return a.mm.BeginBlock(ctx, req)
}

func (a *App) EndBlocker(ctx sdk.Context, req abcitypes.RequestEndBlock) abcitypes.ResponseEndBlock {
	return a.mm.EndBlock(ctx, req)
}

func (a *App) InitChainer(ctx sdk.Context, req abcitypes.RequestInitChain) abcitypes.ResponseInitChain {
	var state map[string]json.RawMessage
	if err := tmjson.Unmarshal(req.AppStateBytes, &state); err != nil {
		panic("failed to unmarshal the app state: " + err.Error())
	}

	a.UpgradeKeeper.SetModuleVersionMap(ctx, a.mm.GetVersionMap())
	return a.mm.InitGenesis(ctx, a.Codec, state)
}

func (a *App) LoadHeight(height int64) error {
	return a.LoadVersion(height)
}

func (a *App) ModuleAccountAddrs() map[string]bool {
	addrs := make(map[string]bool)
	for v := range ModuleAccPerms() {
		addr := authtypes.NewModuleAddress(v)
		addrs[addr.String()] = true
	}

	return addrs
}

func (a *App) SimulationManager() *module.SimulationManager {
	return a.sm
}

func (a *App) RegisterAPIRoutes(server *api.Server, _ serverconfig.APIConfig) {
	ctx := server.ClientCtx
	rpc.RegisterRoutes(ctx, server.Router)
	authrest.RegisterTxRoutes(ctx, server.Router)
	authtx.RegisterGRPCGatewayRoutes(ctx, server.GRPCGatewayRouter)
	tmservice.RegisterGRPCGatewayRoutes(ctx, server.GRPCGatewayRouter)

	ModuleBasics.RegisterRESTRoutes(ctx, server.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(ctx, server.GRPCGatewayRouter)
}

func (a *App) RegisterTxService(ctx client.Context) {
	authtx.RegisterTxService(a.BaseApp.GRPCQueryRouter(), ctx, a.BaseApp.Simulate, a.InterfaceRegistry)
}

func (a *App) RegisterTendermintService(ctx client.Context) {
	tmservice.RegisterTendermintService(a.BaseApp.GRPCQueryRouter(), ctx, a.InterfaceRegistry)
}

func (a *App) SetupAnteHandler(wasmConfig wasmtypes.WasmConfig) {
	handler, err := ante.NewHandler(
		ante.HandlerOptions{
			HandlerOptions: authante.HandlerOptions{
				AccountKeeper:   a.AccountKeeper,
				BankKeeper:      a.BankKeeper,
				FeegrantKeeper:  a.FeeGrantKeeper,
				SignModeHandler: a.TxConfig.SignModeHandler(),
				SigGasConsumer:  authante.DefaultSigVerificationGasConsumer,
			},
			TxCounterStoreKey: a.KV(wasmtypes.StoreKey),
			IBCKeeper:         a.IBCKeeper,
			WasmConfig:        wasmConfig,
		},
	)
	if err != nil {
		panic("failed to create the ante handler: " + err.Error())
	}

	a.SetAnteHandler(handler)
}

func (a *App) SetUpgradeStoreLoader() {
	upgradeInfo, err := a.UpgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic("failed to read the upgrade info from disk: " + err.Error())
	}

	if a.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		return
	}

	if upgradeInfo.Name == upgrades.Name {
		a.SetStoreLoader(
			upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, upgrades.StoreUpgrades),
		)
	}
}

func (a *App) SetUpgradeHandler(configurator module.Configurator) {
	a.UpgradeKeeper.SetUpgradeHandler(
		upgrades.Name,
		upgrades.Handler(
			a.mm, configurator,
			a.IBCICAControllerKeeper,
			a.IBCICAHostKeeper,
		),
	)
}

func (a *App) RegisterSnapshotExtensions() {
	if m := a.SnapshotManager(); m != nil {
		if err := m.RegisterExtensions(
			wasmkeeper.NewWasmSnapshotter(a.CommitMultiStore(), &a.WasmKeeper),
		); err != nil {
			panic("failed to register the snapshot extension: " + err.Error())
		}
	}
}
