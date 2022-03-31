package hub

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
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/server/api"
	serverconfig "github.com/cosmos/cosmos-sdk/server/config"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
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
	"github.com/cosmos/cosmos-sdk/x/authz"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	authzmodule "github.com/cosmos/cosmos-sdk/x/authz/module"
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
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	feegrantmodule "github.com/cosmos/cosmos-sdk/x/feegrant/module"
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
	ibcporttypes "github.com/cosmos/ibc-go/v2/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v2/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v2/modules/core/keeper"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmdb "github.com/tendermint/tm-db"

	hubparams "github.com/sentinel-official/hub/params"
	upgrade3 "github.com/sentinel-official/hub/upgrades/upgrade-3"
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
		authzmodule.AppModuleBasic{},
		bank.AppModuleBasic{},
		capability.AppModule{},
		crisis.AppModuleBasic{},
		distribution.AppModuleBasic{},
		evidence.AppModuleBasic{},
		feegrantmodule.AppModuleBasic{},
		genutil.AppModuleBasic{},
		gov.NewAppModuleBasic(
			distributionclient.ProposalHandler,
			ibcclientclient.UpdateClientProposalHandler,
			ibcclientclient.UpgradeProposalHandler,
			paramsclient.ProposalHandler,
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

type App struct {
	*baseapp.BaseApp

	invarCheckPeriod  uint
	amino             *codec.LegacyAmino
	cdc               codec.Codec
	interfaceRegistry codectypes.InterfaceRegistry

	keys  map[string]*sdk.KVStoreKey
	tkeys map[string]*sdk.TransientStoreKey
	mkeys map[string]*sdk.MemoryStoreKey

	accountKeeper      authkeeper.AccountKeeper
	authzKeeper        authzkeeper.Keeper
	bankKeeper         bankkeeper.Keeper
	capabilityKeeper   *capabilitykeeper.Keeper
	crisisKeeper       crisiskeeper.Keeper
	distributionKeeper distributionkeeper.Keeper
	evidenceKeeper     evidencekeeper.Keeper
	feeGrantKeeper     feegrantkeeper.Keeper
	govKeeper          govkeeper.Keeper
	ibcKeeper          *ibckeeper.Keeper
	ibcTransferKeeper  ibctransferkeeper.Keeper
	mintKeeper         mintkeeper.Keeper
	paramsKeeper       paramskeeper.Keeper
	slashingKeeper     slashingkeeper.Keeper
	stakingKeeper      stakingkeeper.Keeper
	upgradeKeeper      upgradekeeper.Keeper
	customMintKeeper   custommintkeeper.Keeper
	swapKeeper         swapkeeper.Keeper
	vpnKeeper          vpnkeeper.Keeper

	scopedIBCKeeper         capabilitykeeper.ScopedKeeper
	scopedIBCTransferKeeper capabilitykeeper.ScopedKeeper

	configurator      module.Configurator
	moduleManager     *module.Manager
	simulationManager *module.SimulationManager
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
		cdc               = encodingConfig.Marshaler
		amino             = encodingConfig.Amino
		interfaceRegistry = encodingConfig.InterfaceRegistry
		tkeys             = sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
		mkeys             = sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)
		keys              = sdk.NewKVStoreKeys(
			authtypes.StoreKey, authzkeeper.StoreKey, banktypes.StoreKey, capabilitytypes.StoreKey,
			distributiontypes.StoreKey, evidencetypes.StoreKey, feegrant.StoreKey, govtypes.StoreKey,
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
		cdc:               cdc,
		keys:              keys,
		tkeys:             tkeys,
		mkeys:             mkeys,
		interfaceRegistry: interfaceRegistry,
		invarCheckPeriod:  invarCheckPeriod,
	}

	app.configurator = module.NewConfigurator(app.cdc, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.paramsKeeper = paramskeeper.NewKeeper(
		app.cdc,
		app.amino,
		keys[paramstypes.StoreKey],
		tkeys[paramstypes.TStoreKey],
	)

	app.paramsKeeper.Subspace(authtypes.ModuleName)
	app.paramsKeeper.Subspace(banktypes.ModuleName)
	app.paramsKeeper.Subspace(crisistypes.ModuleName)
	app.paramsKeeper.Subspace(distributiontypes.ModuleName)
	app.paramsKeeper.Subspace(evidencetypes.ModuleName)
	app.paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	app.paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	app.paramsKeeper.Subspace(ibchost.ModuleName)
	app.paramsKeeper.Subspace(minttypes.ModuleName)
	app.paramsKeeper.Subspace(slashingtypes.ModuleName)
	app.paramsKeeper.Subspace(stakingtypes.ModuleName)
	app.paramsKeeper.Subspace(swaptypes.ModuleName)

	baseApp.SetParamStore(
		app.paramsKeeper.
			Subspace(baseapp.Paramspace).
			WithKeyTable(paramskeeper.ConsensusParamsKeyTable()),
	)

	app.capabilityKeeper = capabilitykeeper.NewKeeper(
		app.cdc,
		app.keys[capabilitytypes.StoreKey],
		app.mkeys[capabilitytypes.MemStoreKey],
	)

	var (
		scopedIBCKeeper      = app.capabilityKeeper.ScopeToModule(ibchost.ModuleName)
		scopedTransferKeeper = app.capabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	)

	app.accountKeeper = authkeeper.NewAccountKeeper(
		app.cdc,
		app.keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.ModuleAccountsPermissions(),
	)
	app.bankKeeper = bankkeeper.NewBaseKeeper(
		app.cdc,
		app.keys[banktypes.StoreKey],
		app.accountKeeper,
		app.GetSubspace(banktypes.ModuleName),
		app.ModuleAccountAddrs(),
	)
	app.authzKeeper = authzkeeper.NewKeeper(
		app.keys[authzkeeper.StoreKey],
		app.cdc,
		app.BaseApp.MsgServiceRouter(),
	)
	app.feeGrantKeeper = feegrantkeeper.NewKeeper(
		app.cdc,
		app.keys[feegrant.StoreKey],
		app.accountKeeper,
	)
	stakingKeeper := stakingkeeper.NewKeeper(
		app.cdc,
		app.keys[stakingtypes.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		app.GetSubspace(stakingtypes.ModuleName),
	)
	app.mintKeeper = mintkeeper.NewKeeper(
		app.cdc,
		app.keys[minttypes.StoreKey],
		app.GetSubspace(minttypes.ModuleName),
		&stakingKeeper,
		app.accountKeeper,
		app.bankKeeper,
		authtypes.FeeCollectorName,
	)
	app.distributionKeeper = distributionkeeper.NewKeeper(
		app.cdc,
		app.keys[distributiontypes.StoreKey],
		app.GetSubspace(distributiontypes.ModuleName),
		app.accountKeeper,
		app.bankKeeper,
		&stakingKeeper,
		authtypes.FeeCollectorName,
		app.ModuleAccountAddrs(),
	)
	app.slashingKeeper = slashingkeeper.NewKeeper(
		app.cdc,
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
	app.upgradeKeeper = upgradekeeper.NewKeeper(
		skipUpgradeHeights,
		app.keys[upgradetypes.StoreKey],
		app.cdc,
		homePath,
		app.BaseApp,
	)
	app.stakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(
			app.distributionKeeper.Hooks(),
			app.slashingKeeper.Hooks(),
		),
	)
	app.ibcKeeper = ibckeeper.NewKeeper(
		app.cdc,
		app.keys[ibchost.StoreKey],
		app.GetSubspace(ibchost.ModuleName),
		app.stakingKeeper,
		app.upgradeKeeper,
		scopedIBCKeeper,
	)

	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramsproposal.RouterKey, params.NewParamChangeProposalHandler(app.paramsKeeper)).
		AddRoute(distributiontypes.RouterKey, distribution.NewCommunityPoolSpendProposalHandler(app.distributionKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.ibcKeeper.ClientKeeper)).
		AddRoute(ibchost.RouterKey, ibcclient.NewClientProposalHandler(app.ibcKeeper.ClientKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.upgradeKeeper))

	app.govKeeper = govkeeper.NewKeeper(
		app.cdc,
		app.keys[govtypes.StoreKey],
		app.GetSubspace(govtypes.ModuleName),
		app.accountKeeper,
		app.bankKeeper,
		&stakingKeeper,
		govRouter,
	)

	app.ibcTransferKeeper = ibctransferkeeper.NewKeeper(
		app.cdc,
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
		app.cdc,
		app.keys[evidencetypes.StoreKey],
		&app.stakingKeeper,
		app.slashingKeeper,
	)
	app.evidenceKeeper.SetRouter(evidenceRouter)

	app.customMintKeeper = custommintkeeper.NewKeeper(
		app.cdc,
		app.keys[customminttypes.StoreKey],
		app.mintKeeper,
	)
	app.swapKeeper = swapkeeper.NewKeeper(
		app.cdc,
		app.keys[swaptypes.StoreKey],
		app.GetSubspace(swaptypes.ModuleName),
		app.accountKeeper,
		app.bankKeeper,
	)
	app.vpnKeeper = vpnkeeper.NewKeeper(
		app.cdc,
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

	app.moduleManager = module.NewManager(
		auth.NewAppModule(app.cdc, app.accountKeeper, nil),
		authvesting.NewAppModule(app.accountKeeper, app.bankKeeper),
		authzmodule.NewAppModule(app.cdc, app.authzKeeper, app.accountKeeper, app.bankKeeper, app.interfaceRegistry),
		bank.NewAppModule(app.cdc, app.bankKeeper, app.accountKeeper),
		capability.NewAppModule(app.cdc, *app.capabilityKeeper),
		crisis.NewAppModule(&app.crisisKeeper, skipGenesisInvariants),
		distribution.NewAppModule(app.cdc, app.distributionKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		evidence.NewAppModule(app.evidenceKeeper),
		feegrantmodule.NewAppModule(app.cdc, app.accountKeeper, app.bankKeeper, app.feeGrantKeeper, app.interfaceRegistry),
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx, encodingConfig.TxConfig),
		gov.NewAppModule(app.cdc, app.govKeeper, app.accountKeeper, app.bankKeeper),
		ibc.NewAppModule(app.ibcKeeper),
		params.NewAppModule(app.paramsKeeper),
		mint.NewAppModule(app.cdc, app.mintKeeper, app.accountKeeper),
		slashing.NewAppModule(app.cdc, app.slashingKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		staking.NewAppModule(app.cdc, app.stakingKeeper, app.accountKeeper, app.bankKeeper),
		upgrade.NewAppModule(app.upgradeKeeper),
		transferModule,
		custommint.NewAppModule(cdc, app.customMintKeeper),
		swap.NewAppModule(app.cdc, app.swapKeeper),
		vpn.NewAppModule(app.cdc, app.accountKeeper, app.bankKeeper, app.vpnKeeper),
	)

	// NOTE: order is very important here
	app.moduleManager.SetOrderBeginBlockers(
		upgradetypes.ModuleName, customminttypes.ModuleName, minttypes.ModuleName, distributiontypes.ModuleName,
		slashingtypes.ModuleName, evidencetypes.ModuleName, stakingtypes.ModuleName,
		ibchost.ModuleName,
	)
	app.moduleManager.SetOrderEndBlockers(
		crisistypes.ModuleName, govtypes.ModuleName, stakingtypes.ModuleName,
		feegrant.ModuleName, authz.ModuleName, vpntypes.ModuleName,
	)
	app.moduleManager.SetOrderInitGenesis(
		capabilitytypes.ModuleName, authtypes.ModuleName, banktypes.ModuleName,
		distributiontypes.ModuleName, stakingtypes.ModuleName, slashingtypes.ModuleName,
		govtypes.ModuleName, minttypes.ModuleName, crisistypes.ModuleName,
		ibchost.ModuleName, genutiltypes.ModuleName, evidencetypes.ModuleName,
		ibctransfertypes.ModuleName, feegrant.ModuleName, authz.ModuleName,
		customminttypes.ModuleName, swaptypes.ModuleName, vpntypes.ModuleName,
	)

	app.moduleManager.RegisterInvariants(&app.crisisKeeper)
	app.moduleManager.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)
	app.moduleManager.RegisterServices(app.configurator)

	app.simulationManager = module.NewSimulationManager(
		auth.NewAppModule(app.cdc, app.accountKeeper, authsimulation.RandomGenesisAccounts),
		authzmodule.NewAppModule(app.cdc, app.authzKeeper, app.accountKeeper, app.bankKeeper, app.interfaceRegistry),
		bank.NewAppModule(app.cdc, app.bankKeeper, app.accountKeeper),
		capability.NewAppModule(app.cdc, *app.capabilityKeeper),
		distribution.NewAppModule(app.cdc, app.distributionKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		evidence.NewAppModule(app.evidenceKeeper),
		feegrantmodule.NewAppModule(app.cdc, app.accountKeeper, app.bankKeeper, app.feeGrantKeeper, app.interfaceRegistry),
		gov.NewAppModule(app.cdc, app.govKeeper, app.accountKeeper, app.bankKeeper),
		ibc.NewAppModule(app.ibcKeeper),
		mint.NewAppModule(app.cdc, app.mintKeeper, app.accountKeeper),
		params.NewAppModule(app.paramsKeeper),
		slashing.NewAppModule(app.cdc, app.slashingKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		staking.NewAppModule(app.cdc, app.stakingKeeper, app.accountKeeper, app.bankKeeper),
		transferModule,
		custommint.NewAppModule(cdc, app.customMintKeeper),
		swap.NewAppModule(app.cdc, app.swapKeeper),
		vpn.NewAppModule(app.cdc, app.accountKeeper, app.bankKeeper, app.vpnKeeper),
	)
	app.simulationManager.RegisterStoreDecoders()

	app.MountKVStores(app.keys)
	app.MountTransientStores(app.tkeys)
	app.MountMemoryStores(app.mkeys)

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)

	anteHandler, err := NewAnteHandler(
		HandlerOptions{
			HandlerOptions: ante.HandlerOptions{
				AccountKeeper:   app.accountKeeper,
				BankKeeper:      app.bankKeeper,
				FeegrantKeeper:  app.feeGrantKeeper,
				SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
				SigGasConsumer:  ante.DefaultSigVerificationGasConsumer,
			},
			IBCChannelKeeper: app.ibcKeeper.ChannelKeeper,
		},
	)
	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
	app.SetEndBlocker(app.EndBlocker)

	app.upgradeKeeper.SetUpgradeHandler(
		upgrade3.Name,
		upgrade3.Handler(app.moduleManager, app.configurator, &app.ibcKeeper.ConnectionKeeper),
	)

	upgradeInfo, err := app.upgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == upgrade3.Name && !app.upgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{
				authz.ModuleName,
				feegrant.ModuleName,
			},
		}

		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}
	}

	app.scopedIBCKeeper = scopedIBCKeeper
	app.scopedIBCTransferKeeper = scopedTransferKeeper

	return app
}

func (a *App) Name() string { return a.BaseApp.Name() }

func (a *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	if ctx.BlockHeight() == 5_140_000 {
		a.moduleManager.SetOrderBeginBlockers(
			upgradetypes.ModuleName, capabilitytypes.ModuleName, customminttypes.ModuleName,
			minttypes.ModuleName, distributiontypes.ModuleName, slashingtypes.ModuleName,
			evidencetypes.ModuleName, stakingtypes.ModuleName, ibchost.ModuleName,
		)
	}
	return a.moduleManager.BeginBlock(ctx, req)
}

func (a *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return a.moduleManager.EndBlock(ctx, req)
}

func (a *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var state map[string]json.RawMessage
	if err := tmjson.Unmarshal(req.AppStateBytes, &state); err != nil {
		panic(err)
	}

	a.upgradeKeeper.SetModuleVersionMap(ctx, a.moduleManager.GetVersionMap())

	return a.moduleManager.InitGenesis(ctx, a.cdc, state)
}

func (a *App) LegacyAmino() *codec.LegacyAmino {
	return a.amino
}

func (a *App) AppCodec() codec.Codec {
	return a.cdc
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
