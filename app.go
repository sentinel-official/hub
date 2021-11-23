package hub

import (
	"encoding/json"
	"io"
	"os"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/grpc/tmservice"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsimulation "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authvesting "github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/capability"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	distributionclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
	distributionkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/cosmos/cosmos-sdk/x/mint"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramsproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibctransfer "github.com/cosmos/ibc-go/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/modules/core"
	ibcclient "github.com/cosmos/ibc-go/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/modules/core/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	hubparams "github.com/sentinel-official/hub/params"
	upgrade1 "github.com/sentinel-official/hub/upgrades/upgrade-1"
	upgrade2 "github.com/sentinel-official/hub/upgrades/upgrade-2"
	deposittypes "github.com/sentinel-official/hub/x/deposit/types"
	custommint "github.com/sentinel-official/hub/x/mint"
	custommintkeeper "github.com/sentinel-official/hub/x/mint/keeper"
	customminttypes "github.com/sentinel-official/hub/x/mint/types"
	"github.com/sentinel-official/hub/x/swap"
	swapkeeper "github.com/sentinel-official/hub/x/swap/keeper"
	swaptypes "github.com/sentinel-official/hub/x/swap/types"
	"github.com/sentinel-official/hub/x/vpn"
	vpnkeeper "github.com/sentinel-official/hub/x/vpn/keeper"
	vpntypes "github.com/sentinel-official/hub/x/vpn/types"
)

const (
	appName = "Sentinel Hub"
)

var (
	DefaultNodeHome = os.ExpandEnv("${HOME}/.sentinelhub")
	ModuleBasics    = module.NewBasicManager(
		auth.AppModuleBasic{},
		authvesting.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModule{},
		crisis.AppModuleBasic{},
		distribution.AppModuleBasic{},
		evidence.AppModuleBasic{},
		genutil.AppModuleBasic{},
		gov.NewAppModuleBasic(
			distributionclient.ProposalHandler,
			paramsclient.ProposalHandler,
			ibcclientclient.UpdateClientProposalHandler, ibcclientclient.UpgradeProposalHandler,
			upgradeclient.ProposalHandler,
			upgradeclient.CancelProposalHandler,
		),
		ibc.AppModuleBasic{},
		mint.AppModuleBasic{},
		params.AppModuleBasic{},
		slashing.AppModuleBasic{},
		staking.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		custommint.AppModule{},
		swap.AppModuleBasic{},
		vpn.AppModuleBasic{},
	)
)

var (
	_ simapp.App              = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)

// GaiaApp extends an ABCI application, but with most of its parameters exported.
// They are exported for convenience in creating helper functions, as object
// capabilities aren't needed for testing.
type App struct { // nolint: golint
	*baseapp.BaseApp
	legacyAmino       *codec.LegacyAmino
	appCodec          codec.Codec
	interfaceRegistry types.InterfaceRegistry

	invCheckPeriod uint

	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	DistrKeeper      distributionkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidencekeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper
	//	FeeGrantKeeper   feegrantkeeper.Keeper
	//	AuthzKeeper      authzkeeper.Keeper
	//	RouterKeeper     routerkeeper.Keeper
	swapKeeper swapkeeper.Keeper
	vpnKeeper  vpnkeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper

	// the module manager
	mm *module.Manager

	// simulation manager
	sm           *module.SimulationManager
	configurator module.Configurator
}

func NewApp(
	logger log.Logger,
	db tmdb.DB,
	tracer io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	homePath string,
	invarCheckPeriod uint,
	encodingConfig hubparams.EncodingConfig,
	appOptions servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	var (
		appCodec          = encodingConfig.Marshaler
		amino             = encodingConfig.Amino
		interfaceRegistry = encodingConfig.InterfaceRegistry
		tkeys             = sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
		mkeys             = sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)
		keys              = sdk.NewKVStoreKeys(
			authtypes.StoreKey, banktypes.StoreKey, capabilitytypes.StoreKey,
			distributiontypes.StoreKey, evidencetypes.StoreKey, govtypes.StoreKey,
			ibchost.StoreKey, ibctransfertypes.StoreKey, minttypes.StoreKey,
			paramstypes.StoreKey, slashingtypes.StoreKey, stakingtypes.StoreKey,
			upgradetypes.StoreKey, customminttypes.StoreKey, swaptypes.StoreKey, vpntypes.StoreKey,
		)
	)

	baseApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	baseApp.SetCommitMultiStoreTracer(tracer)
	baseApp.SetVersion(version.Version)
	baseApp.SetInterfaceRegistry(interfaceRegistry)

	app := &App{
		BaseApp:           baseApp,
		amino:             amino,
		appCodec:          appCodec,
		keys:              keys,
		tkeys:             tkeys,
		mkeys:             mkeys,
		interfaceRegistry: interfaceRegistry,
		invarCheckPeriod:  invarCheckPeriod,
	}

	app.ParamsKeeper = paramskeeper.NewKeeper(
		app.appCodec,
		app.amino,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	app.ParamsKeeper.Subspace(authtypes.ModuleName)
	app.ParamsKeeper.Subspace(banktypes.ModuleName)
	app.ParamsKeeper.Subspace(crisistypes.ModuleName)
	app.ParamsKeeper.Subspace(distributiontypes.ModuleName)
	app.ParamsKeeper.Subspace(evidencetypes.ModuleName)
	app.ParamsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	app.ParamsKeeper.Subspace(ibctransfertypes.ModuleName)
	app.ParamsKeeper.Subspace(ibchost.ModuleName)
	app.ParamsKeeper.Subspace(minttypes.ModuleName)
	app.ParamsKeeper.Subspace(slashingtypes.ModuleName)
	app.ParamsKeeper.Subspace(stakingtypes.ModuleName)
	app.ParamsKeeper.Subspace(swaptypes.ModuleName)

	baseApp.SetParamStore(
		app.ParamsKeeper.
			Subspace(baseapp.Paramspace).
			WithKeyTable(paramskeeper.ConsensusParamsKeyTable()),
	)

	app.CapabilityKeeper = capabilitykeeper.NewKeeper(
		app.appCodec,
		app.keys[capabilitytypes.StoreKey],
		app.mkeys[capabilitytypes.MemStoreKey],
	)

	var (
		scopedIBCKeeper      = app.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
		scopedTransferKeeper = app.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	)

	app.AccountKeeper = authkeeper.NewAccountKeeper(
		app.appCodec,
		app.keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.ModuleAccountsPermissions(),
	)
	app.BankKeeper = bankkeeper.NewBaseKeeper(
		app.appCodec,
		app.keys[banktypes.StoreKey],
		app.AccountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.ModuleAccountAddrs(),
	)
	stakingKeeper := stakingkeeper.NewKeeper(
		app.appCodec,
		app.keys[stakingtypes.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)
	app.mintKeeper = mintkeeper.NewKeeper(
		app.appCodec,
		app.keys[minttypes.StoreKey],
		app.GetSubspace(minttypes.ModuleName),
		&stakingKeeper,
		app.accountKeeper,
		app.bankKeeper,
		authtypes.FeeCollectorName,
	)
	app.distributionKeeper = distributionkeeper.NewKeeper(
		app.appCodec,
		app.keys[distributiontypes.StoreKey],
		app.GetSubspace(distributiontypes.ModuleName),
		app.accountKeeper,
		app.bankKeeper,
		&stakingKeeper,
		authtypes.FeeCollectorName,
		app.ModuleAccountAddrs(),
	)
	app.slashingKeeper = slashingkeeper.NewKeeper(
		app.appCodec,
		app.keys[slashingtypes.StoreKey],
		&stakingKeeper,
		app.GetSubspace(slashingtypes.ModuleName),
	)
	app.crisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName),
		app.invarCheckPeriod,
		app.bankKeeper,
		authtypes.FeeCollectorName,
	)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
	)
	app.stakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(
			app.distributionKeeper.Hooks(),
			app.slashingKeeper.Hooks(),
		),
	)

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey], app.GetSubspace(ibchost.ModuleName), app.StakingKeeper, app.UpgradeKeeper, scopedIBCKeeper,
	)

	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramsproposal.RouterKey, params.NewParamChangeProposalHandler(app.paramsKeeper)).
		AddRoute(distributiontypes.RouterKey, distribution.NewCommunityPoolSpendProposalHandler(app.distributionKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.upgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper))

	app.govKeeper = govkeeper.NewKeeper(
		app.appCodec,
		app.keys[govtypes.StoreKey],
		app.GetSubspace(govtypes.ModuleName),
		app.accountKeeper,
		app.bankKeeper,
		&stakingKeeper,
		govRouter,
	)

	app.ibcTransferKeeper = ibctransferkeeper.NewKeeper(
		app.appCodec,
		app.keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.ibcKeeper.ChannelKeeper,
		&app.ibcKeeper.PortKeeper,
		app.accountKeeper,
		app.bankKeeper,
		scopedTransferKeeper,
	)

	var (
		evidenceRouter = evidencetypes.NewRouter()
		ibcRouter      = ibcporttypes.NewRouter()
		transferModule = ibctransfer.NewAppModule(app.ibcTransferKeeper)
	)

	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferModule)
	app.ibcKeeper.SetRouter(ibcRouter)

	app.evidenceKeeper = *evidencekeeper.NewKeeper(
		app.appCodec,
		app.keys[evidencetypes.StoreKey],
		&app.stakingKeeper,
		app.slashingKeeper,
	)
	app.evidenceKeeper.SetRouter(evidenceRouter)

	app.customMintKeeper = custommintkeeper.NewKeeper(
		app.appCodec,
		app.keys[customminttypes.StoreKey],
		app.mintKeeper,
	)
	app.swapKeeper = swapkeeper.NewKeeper(
		app.appCodec,
		app.keys[swaptypes.StoreKey],
		app.GetSubspace(swaptypes.ModuleName),
		app.accountKeeper,
		app.bankKeeper,
	)
	app.vpnKeeper = vpnkeeper.NewKeeper(
		app.appCodec,
		app.keys[vpntypes.StoreKey],
		app.paramsKeeper,
		app.accountKeeper,
		app.bankKeeper,
		app.distributionKeeper,
	)

	var (
		skipGenesisInvariants = false
		opt                   = appOptions.Get(crisis.FlagSkipGenesisInvariants)
	)
	if opt, ok := opt.(bool); ok {
		skipGenesisInvariants = opt
	}

	app.manager = module.NewManager(
		auth.NewAppModule(app.appCodec, app.accountKeeper, nil),
		authvesting.NewAppModule(app.accountKeeper, app.bankKeeper),
		bank.NewAppModule(app.appCodec, app.bankKeeper, app.accountKeeper),
		capability.NewAppModule(app.appCodec, *app.capabilityKeeper),
		crisis.NewAppModule(&app.crisisKeeper, skipGenesisInvariants),
		distribution.NewAppModule(app.appCodec, app.distributionKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		evidence.NewAppModule(app.evidenceKeeper),
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx, encodingConfig.TxConfig),
		gov.NewAppModule(app.appCodec, app.govKeeper, app.accountKeeper, app.bankKeeper),
		ibc.NewAppModule(app.ibcKeeper),
		params.NewAppModule(app.paramsKeeper),
		mint.NewAppModule(app.appCodec, app.mintKeeper, app.accountKeeper),
		slashing.NewAppModule(app.appCodec, app.slashingKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		staking.NewAppModule(app.appCodec, app.stakingKeeper, app.accountKeeper, app.bankKeeper),
		upgrade.NewAppModule(app.upgradeKeeper),
		transferModule,
		custommint.NewAppModule(appCodec, app.customMintKeeper),
		swap.NewAppModule(app.appCodec, app.swapKeeper),
		vpn.NewAppModule(app.appCodec, app.accountKeeper, app.bankKeeper, app.vpnKeeper),
	)

	// NOTE: order is very important here
	app.manager.SetOrderBeginBlockers(
		upgradetypes.ModuleName, customminttypes.ModuleName, minttypes.ModuleName, distributiontypes.ModuleName,
		slashingtypes.ModuleName, evidencetypes.ModuleName, stakingtypes.ModuleName,
		ibchost.ModuleName,
	)
	app.manager.SetOrderEndBlockers(
		crisistypes.ModuleName, govtypes.ModuleName, stakingtypes.ModuleName,
		vpntypes.ModuleName,
	)
	app.manager.SetOrderInitGenesis(
		capabilitytypes.ModuleName, authtypes.ModuleName, banktypes.ModuleName,
		distributiontypes.ModuleName, stakingtypes.ModuleName, slashingtypes.ModuleName,
		govtypes.ModuleName, minttypes.ModuleName, crisistypes.ModuleName,
		ibchost.ModuleName, genutiltypes.ModuleName, evidencetypes.ModuleName,
		ibctransfertypes.ModuleName, customminttypes.ModuleName, swaptypes.ModuleName, vpntypes.ModuleName,
	)

	app.manager.RegisterInvariants(&app.crisisKeeper)
	app.manager.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.manager.RegisterServices(module.NewConfigurator(app.MsgServiceRouter(), app.GRPCQueryRouter()))

	app.simulationManager = module.NewSimulationManager(
		auth.NewAppModule(app.appCodec, app.accountKeeper, authsimulation.RandomGenesisAccounts),
		bank.NewAppModule(app.appCodec, app.bankKeeper, app.accountKeeper),
		capability.NewAppModule(app.appCodec, *app.capabilityKeeper),
		distribution.NewAppModule(app.appCodec, app.distributionKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		evidence.NewAppModule(app.evidenceKeeper),
		gov.NewAppModule(app.appCodec, app.govKeeper, app.accountKeeper, app.bankKeeper),
		ibc.NewAppModule(app.ibcKeeper),
		mint.NewAppModule(app.appCodec, app.mintKeeper, app.accountKeeper),
		params.NewAppModule(app.paramsKeeper),
		slashing.NewAppModule(app.appCodec, app.slashingKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		staking.NewAppModule(app.appCodec, app.stakingKeeper, app.accountKeeper, app.bankKeeper),
		transferModule,
		custommint.NewAppModule(appCodec, app.customMintKeeper),
		swap.NewAppModule(app.appCodec, app.swapKeeper),
		vpn.NewAppModule(app.appCodec, app.accountKeeper, app.bankKeeper, app.vpnKeeper),
	)
	app.simulationManager.RegisterStoreDecoders()

	app.MountKVStores(app.keys)
	app.MountTransientStores(app.tkeys)
	app.MountMemoryStores(app.mkeys)

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(
		ante.NewAnteHandler(
			app.accountKeeper,
			app.bankKeeper,
			ante.DefaultSigVerificationGasConsumer,
			encodingConfig.TxConfig.SignModeHandler(),
		),
	)
	app.SetEndBlocker(app.EndBlocker)

	app.upgradeKeeper.SetUpgradeHandler(
		upgrade1.Name,
		upgrade1.Handler(app.accountKeeper),
	)
	app.upgradeKeeper.SetUpgradeHandler(
		upgrade2.Name,
		upgrade2.Handler(app.SetStoreLoader, app.accountKeeper, app.upgradeKeeper, app.customMintKeeper),
	)

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}

		ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})
		app.capabilityKeeper.InitializeAndSeal(ctx)
	}

	app.scopedIBCKeeper = scopedIBCKeeper
	app.scopedIBCTransferKeeper = scopedTransferKeeper

	return app
}

func (a *App) Name() string { return a.BaseApp.Name() }

func (a *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return a.manager.BeginBlock(ctx, req)
}

func (a *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return a.manager.EndBlock(ctx, req)
}

func (a *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var state map[string]json.RawMessage
	if err := tmjson.Unmarshal(req.AppStateBytes, &state); err != nil {
		panic(err)
	}
	return a.manager.InitGenesis(ctx, a.appCodec, state)
}

func (a *App) LegacyAmino() *codec.LegacyAmino {
	return a.amino
}

func (a *App) AppCodec() codec.Codec {
	return a.appCodec
}

func (a *App) InterfaceRegistry() codectypes.InterfaceRegistry {
	return a.interfaceRegistry
}

func (a *App) RegisterTxService(ctx client.Context) {
	authtx.RegisterTxService(a.BaseApp.GRPCQueryRouter(), ctx, a.BaseApp.Simulate, a.interfaceRegistry)
}

func (a *App) RegisterTendermintService(ctx client.Context) {
	tmservice.RegisterTendermintService(a.BaseApp.GRPCQueryRouter(), ctx, a.interfaceRegistry)
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

func (a *App) LoadHeight(height int64) error {
	return a.LoadVersion(height)
}

func (a *App) ModuleAccountsPermissions() map[string][]string {
	return map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distributiontypes.ModuleName:   nil,
		govtypes.ModuleName:            {authtypes.Burner},
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		customminttypes.ModuleName:     nil,
		swaptypes.ModuleName:           {authtypes.Minter},
		deposittypes.ModuleName:        nil,
	}
}

func (a *App) ModuleAccountAddrs() map[string]bool {
	accounts := make(map[string]bool)
	for name := range a.ModuleAccountsPermissions() {
		accounts[authtypes.NewModuleAddress(name).String()] = true
	}

	return accounts
}

func (a *App) SimulationManager() *module.SimulationManager {
	return a.simulationManager
}

func (a *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := a.paramsKeeper.GetSubspace(moduleName)
	return subspace
}
