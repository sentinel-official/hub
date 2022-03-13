package app

import (
	"encoding/json"
	"fmt"
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
	authvestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
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
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
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
	ibctransfer "github.com/cosmos/ibc-go/v2/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v2/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v2/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v2/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v2/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v2/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v2/modules/core/02-client/types"
	ibchost "github.com/cosmos/ibc-go/v2/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v2/modules/core/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmdb "github.com/tendermint/tm-db"

	hubparams "github.com/sentinel-official/hub/params"
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
	FeeGrantKeeper   feegrantkeeper.Keeper
	AuthzKeeper      authzkeeper.Keeper

	//	RouterKeeper     routerkeeper.Keeper
	SwapKeeper       swapkeeper.Keeper
	VpnKeeper        vpnkeeper.Keeper
	CustomMintKeeper custommintkeeper.Keeper

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
	invCheckPeriod uint,
	encodingConfig hubparams.EncodingConfig,
	appOptions servertypes.AppOptions,
	baseAppOptions ...func(*baseapp.BaseApp),
) *App {
	var (
		appCodec          = encodingConfig.Marshaler
		legacyAmino       = encodingConfig.Amino
		interfaceRegistry = encodingConfig.InterfaceRegistry

		tkeys   = sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
		memKeys = sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)
		keys    = sdk.NewKVStoreKeys(
			authtypes.StoreKey,
			banktypes.StoreKey,
			capabilitytypes.StoreKey,
			distributiontypes.StoreKey,
			evidencetypes.StoreKey,
			govtypes.StoreKey,
			ibchost.StoreKey,
			ibctransfertypes.StoreKey,
			minttypes.StoreKey,
			paramstypes.StoreKey,
			slashingtypes.StoreKey,
			stakingtypes.StoreKey,
			upgradetypes.StoreKey,
			customminttypes.StoreKey,
			swaptypes.StoreKey,
			vpntypes.StoreKey,
		)
	)

	baseApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOptions...)
	baseApp.SetCommitMultiStoreTracer(tracer)
	baseApp.SetVersion(version.Version)
	baseApp.SetInterfaceRegistry(interfaceRegistry)

	app := &App{
		BaseApp:           baseApp,
		legacyAmino:       legacyAmino,
		appCodec:          appCodec,
		keys:              keys,
		tkeys:             tkeys,
		memKeys:           memKeys,
		interfaceRegistry: interfaceRegistry,
		invCheckPeriod:    invCheckPeriod,
	}

	app.ParamsKeeper = paramskeeper.NewKeeper(
		app.appCodec,
		app.legacyAmino,
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
		app.memKeys[capabilitytypes.MemStoreKey],
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
		app.AccountKeeper,
		app.BankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)
	app.MintKeeper = mintkeeper.NewKeeper(
		app.appCodec,
		app.keys[minttypes.StoreKey],
		app.GetSubspace(minttypes.ModuleName),
		&stakingKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)
	app.DistrKeeper = distributionkeeper.NewKeeper(
		app.appCodec,
		app.keys[distributiontypes.StoreKey],
		app.GetSubspace(distributiontypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&stakingKeeper,
		authtypes.FeeCollectorName,
		app.ModuleAccountAddrs(),
	)
	app.SlashingKeeper = slashingkeeper.NewKeeper(
		app.appCodec,
		app.keys[slashingtypes.StoreKey],
		&stakingKeeper,
		app.GetSubspace(slashingtypes.ModuleName),
	)
	app.CrisisKeeper = crisiskeeper.NewKeeper(
		app.GetSubspace(crisistypes.ModuleName),
		app.invCheckPeriod,
		app.BankKeeper,
		authtypes.FeeCollectorName,
	)
	app.UpgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		keys[upgradetypes.StoreKey],
		appCodec,
		homePath,
		app.BaseApp,
	)
	app.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(
			app.DistrKeeper.Hooks(),
			app.SlashingKeeper.Hooks(),
		),
	)

	// Create IBC Keeper
	app.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey], app.GetSubspace(ibchost.ModuleName), app.StakingKeeper, app.UpgradeKeeper, scopedIBCKeeper,
	)

	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramsproposal.RouterKey, params.NewParamChangeProposalHandler(app.ParamsKeeper)).
		AddRoute(distributiontypes.RouterKey, distribution.NewCommunityPoolSpendProposalHandler(app.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.IBCKeeper.ClientKeeper))

	app.GovKeeper = govkeeper.NewKeeper(
		app.appCodec,
		app.keys[govtypes.StoreKey],
		app.GetSubspace(govtypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
		&stakingKeeper,
		govRouter,
	)

	app.TransferKeeper = ibctransferkeeper.NewKeeper(
		app.appCodec,
		app.keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.IBCKeeper.ChannelKeeper,
		&app.IBCKeeper.PortKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		scopedTransferKeeper,
	)

	var (
		evidenceRouter = evidencetypes.NewRouter()
		transferModule = ibctransfer.NewAppModule(app.TransferKeeper)
	)

	app.EvidenceKeeper = *evidencekeeper.NewKeeper(
		app.appCodec,
		app.keys[evidencetypes.StoreKey],
		&app.StakingKeeper,
		app.SlashingKeeper,
	)
	app.EvidenceKeeper.SetRouter(evidenceRouter)

	app.CustomMintKeeper = custommintkeeper.NewKeeper(
		app.appCodec,
		app.keys[customminttypes.StoreKey],
		app.MintKeeper,
	)
	app.SwapKeeper = swapkeeper.NewKeeper(
		app.appCodec,
		app.keys[swaptypes.StoreKey],
		app.GetSubspace(swaptypes.ModuleName),
		app.AccountKeeper,
		app.BankKeeper,
	)
	app.VpnKeeper = vpnkeeper.NewKeeper(
		app.appCodec,
		app.keys[vpntypes.StoreKey],
		app.ParamsKeeper,
		app.AccountKeeper,
		app.BankKeeper,
		app.DistrKeeper,
	)

	var (
		skipGenesisInvariants = false
		opt                   = appOptions.Get(crisis.FlagSkipGenesisInvariants)
	)
	if opt, ok := opt.(bool); ok {
		skipGenesisInvariants = opt
	}

	app.mm = module.NewManager(
		auth.NewAppModule(app.appCodec, app.AccountKeeper, nil),
		authvesting.NewAppModule(app.AccountKeeper, app.BankKeeper),
		bank.NewAppModule(app.appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(app.appCodec, *app.CapabilityKeeper),
		crisis.NewAppModule(&app.CrisisKeeper, skipGenesisInvariants),
		distribution.NewAppModule(app.appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		genutil.NewAppModule(app.AccountKeeper, app.StakingKeeper, app.BaseApp.DeliverTx, encodingConfig.TxConfig),
		gov.NewAppModule(app.appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		params.NewAppModule(app.ParamsKeeper),
		mint.NewAppModule(app.appCodec, app.MintKeeper, app.AccountKeeper),
		slashing.NewAppModule(app.appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(app.appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		upgrade.NewAppModule(app.UpgradeKeeper),
		transferModule,
		custommint.NewAppModule(appCodec, app.CustomMintKeeper),
		swap.NewAppModule(app.appCodec, app.SwapKeeper),
		vpn.NewAppModule(app.appCodec, app.AccountKeeper, app.BankKeeper, app.VpnKeeper),
	)

	// NOTE: order is very important here
	app.mm.SetOrderBeginBlockers(
		authvestingtypes.ModuleName, upgradetypes.ModuleName, paramstypes.ModuleName, capabilitytypes.ModuleName, authtypes.ModuleName, banktypes.ModuleName,
		distributiontypes.ModuleName, stakingtypes.ModuleName, slashingtypes.ModuleName,
		govtypes.ModuleName, minttypes.ModuleName, crisistypes.ModuleName,
		ibchost.ModuleName, genutiltypes.ModuleName, evidencetypes.ModuleName,
		ibctransfertypes.ModuleName, customminttypes.ModuleName, swaptypes.ModuleName, vpntypes.ModuleName,
	)
	app.mm.SetOrderEndBlockers(
		authvestingtypes.ModuleName, upgradetypes.ModuleName, paramstypes.ModuleName, capabilitytypes.ModuleName, authtypes.ModuleName, banktypes.ModuleName,
		distributiontypes.ModuleName, stakingtypes.ModuleName, slashingtypes.ModuleName,
		govtypes.ModuleName, minttypes.ModuleName, crisistypes.ModuleName,
		ibchost.ModuleName, genutiltypes.ModuleName, evidencetypes.ModuleName,
		ibctransfertypes.ModuleName, customminttypes.ModuleName, swaptypes.ModuleName, vpntypes.ModuleName,
	)
	app.mm.SetOrderInitGenesis(
		authvestingtypes.ModuleName, upgradetypes.ModuleName, paramstypes.ModuleName, capabilitytypes.ModuleName, authtypes.ModuleName, banktypes.ModuleName,
		distributiontypes.ModuleName, stakingtypes.ModuleName, slashingtypes.ModuleName,
		govtypes.ModuleName, minttypes.ModuleName, crisistypes.ModuleName,
		ibchost.ModuleName, genutiltypes.ModuleName, evidencetypes.ModuleName,
		ibctransfertypes.ModuleName, customminttypes.ModuleName, swaptypes.ModuleName, vpntypes.ModuleName,
	)

	app.mm.RegisterInvariants(&app.CrisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)

	app.configurator = module.NewConfigurator(app.appCodec, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.mm.RegisterServices(app.configurator)

	app.sm = module.NewSimulationManager(
		auth.NewAppModule(app.appCodec, app.AccountKeeper, authsimulation.RandomGenesisAccounts),
		bank.NewAppModule(app.appCodec, app.BankKeeper, app.AccountKeeper),
		capability.NewAppModule(app.appCodec, *app.CapabilityKeeper),
		distribution.NewAppModule(app.appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		evidence.NewAppModule(app.EvidenceKeeper),
		gov.NewAppModule(app.appCodec, app.GovKeeper, app.AccountKeeper, app.BankKeeper),
		ibc.NewAppModule(app.IBCKeeper),
		mint.NewAppModule(app.appCodec, app.MintKeeper, app.AccountKeeper),
		params.NewAppModule(app.ParamsKeeper),
		slashing.NewAppModule(app.appCodec, app.SlashingKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
		staking.NewAppModule(app.appCodec, app.StakingKeeper, app.AccountKeeper, app.BankKeeper),
		transferModule,
		custommint.NewAppModule(appCodec, app.CustomMintKeeper),
		swap.NewAppModule(app.appCodec, app.SwapKeeper),
		vpn.NewAppModule(app.appCodec, app.AccountKeeper, app.BankKeeper, app.VpnKeeper),
	)
	app.sm.RegisterStoreDecoders()

	// initialize stores
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)
	app.MountMemoryStores(memKeys)

	anteHandler, err := NewAnteHandler(
		HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.AccountKeeper,
				BankKeeper:      app.BankKeeper,
				FeegrantKeeper:  app.FeeGrantKeeper,
				SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			IBCChannelkeeper: app.IBCKeeper.ChannelKeeper,
		},
	)
	if err != nil {
		panic(fmt.Errorf("failed to create AnteHandler: %s", err))
	}

	app.SetAnteHandler(anteHandler)
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	/*
		app.UpgradeKeeper.SetUpgradeHandler(
			upgrade1.Name,
			upgrade1.Handler(app.AccountKeeper),
		)
		app.UpgradeKeeper.SetUpgradeHandler(
			upgrade2.Name,
			upgrade2.Handler(app.SetStoreLoader, app.AccountKeeper, app.UpgradeKeeper, app.customMintKeeper),
		)
	*/

	/*
		upgradeInfo, err := app.UpgradeKeeper.ReadUpgradeInfoFromDisk()
		if err != nil {
			panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
		}

		if upgradeInfo.Name == upgradeName && !app.UpgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
			storeUpgrades := store.StoreUpgrades{
				Added: []string{authz.ModuleName, feegrant.ModuleName, routertypes.ModuleName},
			}

			// configure store loader that checks if version == upgradeHeight and applies store upgrades
			app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
		}
	*/

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(fmt.Sprintf("failed to load latest version: %s", err))
		}
	}

	app.ScopedIBCKeeper = scopedIBCKeeper
	app.ScopedTransferKeeper = scopedTransferKeeper

	return app
}

func (a *App) Name() string { return a.BaseApp.Name() }

func (a *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return a.mm.BeginBlock(ctx, req)
}

func (a *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return a.mm.EndBlock(ctx, req)
}

func (a *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var state map[string]json.RawMessage
	if err := tmjson.Unmarshal(req.AppStateBytes, &state); err != nil {
		panic(err)
	}

	fmt.Printf("state = %v \n", state)

	return a.mm.InitGenesis(ctx, a.appCodec, state)
}

func (a *App) LegacyAmino() *codec.LegacyAmino {
	return a.legacyAmino
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
	return a.SimulationManager()
}

func (a *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := a.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}
