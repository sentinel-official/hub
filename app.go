package hub

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/CosmWasm/wasmd/x/wasm"
	wasmclient "github.com/CosmWasm/wasmd/x/wasm/client"
	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
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
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"
	authrest "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authsimulation "github.com/cosmos/cosmos-sdk/x/auth/simulation"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	authvesting "github.com/cosmos/cosmos-sdk/x/auth/vesting"
	authvestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
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
	ibcica "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts"
	ibcicacontroller "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller"
	ibcicacontrollerkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/keeper"
	ibcicacontrollertypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	ibcicahost "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host"
	ibcicahostkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/keeper"
	ibcicahosttypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/types"
	ibcicatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
	ibcfee "github.com/cosmos/ibc-go/v4/modules/apps/29-fee"
	ibcfeekeeper "github.com/cosmos/ibc-go/v4/modules/apps/29-fee/keeper"
	ibcfeetypes "github.com/cosmos/ibc-go/v4/modules/apps/29-fee/types"
	ibctransfer "github.com/cosmos/ibc-go/v4/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v4/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	ibc "github.com/cosmos/ibc-go/v4/modules/core"
	ibcclient "github.com/cosmos/ibc-go/v4/modules/core/02-client"
	ibcclientclient "github.com/cosmos/ibc-go/v4/modules/core/02-client/client"
	ibcclienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"
	"github.com/spf13/cast"
	abci "github.com/tendermint/tendermint/abci/types"
	tmjson "github.com/tendermint/tendermint/libs/json"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"

	hubparams "github.com/sentinel-official/hub/params"
	hubupgrades "github.com/sentinel-official/hub/upgrades"
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
			append(
				wasmclient.ProposalHandlers,
				distributionclient.ProposalHandler,
				ibcclientclient.UpdateClientProposalHandler,
				ibcclientclient.UpgradeProposalHandler,
				paramsclient.ProposalHandler,
				upgradeclient.ProposalHandler,
				upgradeclient.CancelProposalHandler,
			)...,
		),
		ibc.AppModuleBasic{},
		ibcfee.AppModuleBasic{},
		ibcica.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
		mint.AppModuleBasic{},
		params.AppModuleBasic{},
		slashing.AppModuleBasic{},
		staking.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		custommint.AppModule{},
		swap.AppModuleBasic{},
		vpn.AppModuleBasic{},
		wasm.AppModuleBasic{},
	)
	WasmEnableSpecificProposals = ""
	WasmProposalsEnabled        = "true"
)

var (
	_ simapp.App              = (*App)(nil)
	_ servertypes.Application = (*App)(nil)
)

func GetWasmEnabledProposals() []wasmtypes.ProposalType {
	if WasmEnableSpecificProposals == "" {
		if WasmProposalsEnabled == "true" {
			return wasmtypes.EnableAllProposals
		}

		return wasmtypes.DisableAllProposals
	}

	chunks := strings.Split(WasmEnableSpecificProposals, ",")

	proposals, err := wasmtypes.ConvertToProposals(chunks)
	if err != nil {
		panic(err)
	}

	return proposals
}

type App struct {
	*baseapp.BaseApp

	invCheckPeriod    uint
	amino             *codec.LegacyAmino
	cdc               codec.Codec
	interfaceRegistry codectypes.InterfaceRegistry
	keys              map[string]*sdk.KVStoreKey
	tkeys             map[string]*sdk.TransientStoreKey
	mkeys             map[string]*sdk.MemoryStoreKey

	accountKeeper          authkeeper.AccountKeeper
	authzKeeper            authzkeeper.Keeper
	bankKeeper             bankkeeper.Keeper
	capabilityKeeper       *capabilitykeeper.Keeper
	crisisKeeper           crisiskeeper.Keeper
	distributionKeeper     distributionkeeper.Keeper
	evidenceKeeper         evidencekeeper.Keeper
	feeGrantKeeper         feegrantkeeper.Keeper
	govKeeper              govkeeper.Keeper
	ibcKeeper              *ibckeeper.Keeper
	ibcFeeKeeper           ibcfeekeeper.Keeper
	ibcICAControllerKeeper ibcicacontrollerkeeper.Keeper
	ibcICAHostKeeper       ibcicahostkeeper.Keeper
	ibcTransferKeeper      ibctransferkeeper.Keeper
	mintKeeper             mintkeeper.Keeper
	paramsKeeper           paramskeeper.Keeper
	slashingKeeper         slashingkeeper.Keeper
	stakingKeeper          stakingkeeper.Keeper
	upgradeKeeper          upgradekeeper.Keeper
	customMintKeeper       custommintkeeper.Keeper
	swapKeeper             swapkeeper.Keeper
	vpnKeeper              vpnkeeper.Keeper
	wasmKeeper             wasmkeeper.Keeper

	scopedIBCKeeper              capabilitykeeper.ScopedKeeper
	scopedIBCFeeKeeper           capabilitykeeper.ScopedKeeper
	scopedIBCICAControllerKeeper capabilitykeeper.ScopedKeeper
	scopedIBCICAHostKeeper       capabilitykeeper.ScopedKeeper
	scopedIBCTransferKeeper      capabilitykeeper.ScopedKeeper
	scopedWasmKeeper             capabilitykeeper.ScopedKeeper

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
	invCheckPeriod uint,
	encodingConfig hubparams.EncodingConfig,
	enabledProposals []wasmtypes.ProposalType,
	appOpts servertypes.AppOptions,
	wasmOpts []wasmkeeper.Option,
	baseAppOpts ...func(*baseapp.BaseApp),
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
			ibchost.StoreKey, ibcfeetypes.StoreKey, ibcicacontrollertypes.StoreKey, ibcicahosttypes.StoreKey, ibctransfertypes.StoreKey,
			minttypes.StoreKey, paramstypes.StoreKey, slashingtypes.StoreKey, stakingtypes.StoreKey, upgradetypes.StoreKey,
			customminttypes.StoreKey, swaptypes.StoreKey, vpntypes.StoreKey, wasmtypes.StoreKey,
		)
	)

	baseApp := baseapp.NewBaseApp(appName, logger, db, encodingConfig.TxConfig.TxDecoder(), baseAppOpts...)
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
		invCheckPeriod:    invCheckPeriod,
	}

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
	app.paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	app.paramsKeeper.Subspace(ibchost.ModuleName)
	app.paramsKeeper.Subspace(ibcicahosttypes.SubModuleName)
	app.paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	app.paramsKeeper.Subspace(minttypes.ModuleName)
	app.paramsKeeper.Subspace(slashingtypes.ModuleName)
	app.paramsKeeper.Subspace(stakingtypes.ModuleName)
	app.paramsKeeper.Subspace(swaptypes.ModuleName)
	app.paramsKeeper.Subspace(wasmtypes.ModuleName)

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

	app.scopedIBCKeeper = app.capabilityKeeper.ScopeToModule(ibchost.ModuleName)
	app.scopedIBCFeeKeeper = app.capabilityKeeper.ScopeToModule(ibcfeetypes.ModuleName) // TODO: recheck
	app.scopedIBCICAControllerKeeper = app.capabilityKeeper.ScopeToModule(ibcicacontrollertypes.SubModuleName)
	app.scopedIBCICAHostKeeper = app.capabilityKeeper.ScopeToModule(ibcicahosttypes.SubModuleName)
	app.scopedIBCTransferKeeper = app.capabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	app.scopedWasmKeeper = app.capabilityKeeper.ScopeToModule(wasmtypes.ModuleName)
	app.capabilityKeeper.Seal()

	app.accountKeeper = authkeeper.NewAccountKeeper(
		app.cdc,
		app.keys[authtypes.StoreKey],
		app.GetSubspace(authtypes.ModuleName),
		authtypes.ProtoBaseAccount,
		app.ModuleAccountPermissions(),
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
		app.invCheckPeriod,
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
		app.scopedIBCKeeper,
	)
	app.ibcFeeKeeper = ibcfeekeeper.NewKeeper(
		app.cdc,
		app.keys[ibcfeetypes.StoreKey],
		app.GetSubspace(ibcfeetypes.ModuleName),
		app.ibcKeeper.ChannelKeeper,
		app.ibcKeeper.ChannelKeeper,
		&app.ibcKeeper.PortKeeper,
		app.accountKeeper,
		app.bankKeeper,
	)
	app.ibcICAControllerKeeper = ibcicacontrollerkeeper.NewKeeper(
		app.cdc,
		app.keys[ibcicacontrollertypes.StoreKey],
		app.GetSubspace(ibcicacontrollertypes.SubModuleName),
		app.ibcFeeKeeper,
		app.ibcKeeper.ChannelKeeper,
		&app.ibcKeeper.PortKeeper,
		app.scopedIBCICAControllerKeeper,
		app.MsgServiceRouter(),
	)
	app.ibcICAHostKeeper = ibcicahostkeeper.NewKeeper(
		app.cdc,
		app.keys[ibcicahosttypes.StoreKey],
		app.GetSubspace(ibcicahosttypes.SubModuleName),
		app.ibcKeeper.ChannelKeeper,
		&app.ibcKeeper.PortKeeper,
		app.accountKeeper,
		app.scopedIBCICAHostKeeper,
		app.MsgServiceRouter(),
	)
	app.ibcTransferKeeper = ibctransferkeeper.NewKeeper(
		app.cdc,
		app.keys[ibctransfertypes.StoreKey],
		app.GetSubspace(ibctransfertypes.ModuleName),
		app.ibcKeeper.ChannelKeeper,
		app.ibcKeeper.ChannelKeeper,
		&app.ibcKeeper.PortKeeper,
		app.accountKeeper,
		app.bankKeeper,
		app.scopedIBCTransferKeeper,
	)

	app.evidenceKeeper = *evidencekeeper.NewKeeper(
		app.cdc,
		app.keys[evidencetypes.StoreKey],
		&app.stakingKeeper,
		app.slashingKeeper,
	)
	evidenceRouter := evidencetypes.NewRouter()
	app.evidenceKeeper.SetRouter(evidenceRouter) // TODO: recheck

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
		authtypes.FeeCollectorName,
	)

	wasmDir := filepath.Join(homePath, "data")

	wasmConfig, err := wasm.ReadWasmConfig(appOpts)
	if err != nil {
		panic("error while reading wasm config: " + err.Error())
	}

	wasmCapabilities := "iterator,staking,stargate,cosmwasm_1_1" // TODO: add queries?
	app.wasmKeeper = wasmkeeper.NewKeeper(
		app.cdc,
		keys[wasmtypes.StoreKey],
		app.GetSubspace(wasmtypes.ModuleName),
		app.accountKeeper,
		app.bankKeeper,
		app.stakingKeeper,
		app.distributionKeeper,
		app.ibcKeeper.ChannelKeeper,
		&app.ibcKeeper.PortKeeper,
		app.scopedWasmKeeper,
		app.ibcTransferKeeper,
		app.MsgServiceRouter(),
		app.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		wasmCapabilities,
		wasmOpts...,
	)

	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramsproposal.RouterKey, params.NewParamChangeProposalHandler(app.paramsKeeper)).
		AddRoute(distributiontypes.RouterKey, distribution.NewCommunityPoolSpendProposalHandler(app.distributionKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(app.ibcKeeper.ClientKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.upgradeKeeper))

	if len(enabledProposals) != 0 {
		govRouter.AddRoute(wasmtypes.RouterKey, wasmkeeper.NewWasmProposalHandler(app.wasmKeeper, enabledProposals))
	}

	app.govKeeper = govkeeper.NewKeeper(
		app.cdc,
		app.keys[govtypes.StoreKey],
		app.GetSubspace(govtypes.ModuleName),
		app.accountKeeper,
		app.bankKeeper,
		&stakingKeeper,
		govRouter,
	)

	var (
		ibcICAControllerIBCModule ibcporttypes.IBCModule
		ibcICAHostIBCModule       = ibcicahost.NewIBCModule(app.ibcICAHostKeeper)
		ibcTransferIBCModule      = ibctransfer.NewIBCModule(app.ibcTransferKeeper)
		wasmIBCModule             ibcporttypes.IBCModule
	)

	ibcICAControllerIBCModule = ibcicacontroller.NewIBCMiddleware(ibcICAControllerIBCModule, app.ibcICAControllerKeeper)
	ibcICAControllerIBCModule = ibcfee.NewIBCMiddleware(ibcICAControllerIBCModule, app.ibcFeeKeeper)

	wasmIBCModule = wasm.NewIBCHandler(app.wasmKeeper, app.ibcKeeper.ChannelKeeper, app.ibcKeeper.ChannelKeeper) // TODO: recheck
	wasmIBCModule = ibcfee.NewIBCMiddleware(wasmIBCModule, app.ibcFeeKeeper)

	ibcPortRouter := ibcporttypes.NewRouter().
		AddRoute(ibcicacontrollertypes.SubModuleName, ibcICAControllerIBCModule).
		AddRoute(ibcicahosttypes.SubModuleName, ibcICAHostIBCModule).
		AddRoute(ibctransfertypes.ModuleName, ibcTransferIBCModule).
		AddRoute(wasmtypes.ModuleName, wasmIBCModule)
	app.ibcKeeper.SetRouter(ibcPortRouter)

	skipGenesisInvariants := cast.ToBool(appOpts.Get(crisis.FlagSkipGenesisInvariants))

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
		ibcica.NewAppModule(&app.ibcICAControllerKeeper, &app.ibcICAHostKeeper),
		ibcfee.NewAppModule(app.ibcFeeKeeper),
		ibctransfer.NewAppModule(app.ibcTransferKeeper),
		params.NewAppModule(app.paramsKeeper),
		mint.NewAppModule(app.cdc, app.mintKeeper, app.accountKeeper),
		slashing.NewAppModule(app.cdc, app.slashingKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		staking.NewAppModule(app.cdc, app.stakingKeeper, app.accountKeeper, app.bankKeeper),
		upgrade.NewAppModule(app.upgradeKeeper),
		custommint.NewAppModule(cdc, app.customMintKeeper),
		swap.NewAppModule(app.cdc, app.swapKeeper),
		vpn.NewAppModule(app.cdc, app.accountKeeper, app.bankKeeper, app.vpnKeeper),
		wasm.NewAppModule(app.cdc, &app.wasmKeeper, app.stakingKeeper, app.accountKeeper, app.bankKeeper),
	)

	app.moduleManager.SetOrderBeginBlockers(
		upgradetypes.ModuleName,
		capabilitytypes.ModuleName,
		customminttypes.ModuleName,
		minttypes.ModuleName,
		distributiontypes.ModuleName,
		slashingtypes.ModuleName,
		evidencetypes.ModuleName,
		stakingtypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		govtypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		authvestingtypes.ModuleName,
		ibchost.ModuleName,
		ibcicatypes.ModuleName,
		ibcfeetypes.ModuleName,
		ibctransfertypes.ModuleName,
		swaptypes.ModuleName,
		vpntypes.ModuleName,
		wasmtypes.ModuleName,
	)
	app.moduleManager.SetOrderEndBlockers(
		crisistypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distributiontypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		feegrant.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		authvestingtypes.ModuleName,
		ibchost.ModuleName,
		ibcicatypes.ModuleName,
		ibcfeetypes.ModuleName,
		ibctransfertypes.ModuleName,
		customminttypes.ModuleName,
		swaptypes.ModuleName,
		vpntypes.ModuleName,
		wasmtypes.ModuleName,
	)
	app.moduleManager.SetOrderInitGenesis(
		capabilitytypes.ModuleName,
		authtypes.ModuleName,
		banktypes.ModuleName,
		distributiontypes.ModuleName,
		govtypes.ModuleName,
		stakingtypes.ModuleName,
		slashingtypes.ModuleName,
		minttypes.ModuleName,
		crisistypes.ModuleName,
		genutiltypes.ModuleName,
		evidencetypes.ModuleName,
		authz.ModuleName,
		paramstypes.ModuleName,
		upgradetypes.ModuleName,
		authvestingtypes.ModuleName,
		feegrant.ModuleName,
		ibchost.ModuleName,
		ibcicatypes.ModuleName,
		ibcfeetypes.ModuleName,
		ibctransfertypes.ModuleName,
		customminttypes.ModuleName,
		swaptypes.ModuleName,
		vpntypes.ModuleName,
		wasmtypes.ModuleName,
	)

	app.moduleManager.RegisterInvariants(&app.crisisKeeper)
	app.moduleManager.RegisterRoutes(app.Router(), app.QueryRouter(), encodingConfig.Amino)

	app.configurator = module.NewConfigurator(app.cdc, app.MsgServiceRouter(), app.GRPCQueryRouter())
	app.moduleManager.RegisterServices(app.configurator)

	app.MountKVStores(app.keys)
	app.MountTransientStores(app.tkeys)
	app.MountMemoryStores(app.mkeys)

	anteHandler, err := NewAnteHandler(
		HandlerOptions{
			HandlerOptions: authante.HandlerOptions{
				AccountKeeper:   app.accountKeeper,
				BankKeeper:      app.bankKeeper,
				FeegrantKeeper:  app.feeGrantKeeper,
				SignModeHandler: encodingConfig.TxConfig.SignModeHandler(),
				SigGasConsumer:  authante.DefaultSigVerificationGasConsumer,
			},
			TxCounterStoreKey: app.keys[wasmtypes.StoreKey],
			IBCKeeper:         app.ibcKeeper,
			WasmConfig:        wasmConfig,
		},
	)
	if err != nil {
		panic(err)
	}

	app.SetAnteHandler(anteHandler)
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)

	if manager := app.SnapshotManager(); manager != nil {
		err = manager.RegisterExtensions(
			wasmkeeper.NewWasmSnapshotter(app.CommitMultiStore(), &app.wasmKeeper),
		)
		if err != nil {
			panic("failed to register snapshot extension: " + err.Error())
		}
	}

	app.upgradeKeeper.SetUpgradeHandler(
		hubupgrades.Name,
		hubupgrades.Handler(
			app.moduleManager,
			app.configurator,
			app.ibcICAControllerKeeper,
			app.ibcICAHostKeeper,
		),
	)

	upgradeInfo, err := app.upgradeKeeper.ReadUpgradeInfoFromDisk()
	if err != nil {
		panic(fmt.Sprintf("failed to read upgrade info from disk %s", err))
	}

	if upgradeInfo.Name == hubupgrades.Name && !app.upgradeKeeper.IsSkipHeight(upgradeInfo.Height) {
		storeUpgrades := storetypes.StoreUpgrades{
			Added: []string{
				ibcicacontrollertypes.StoreKey,
				ibcfeetypes.StoreKey,
			},
		}

		app.SetStoreLoader(upgradetypes.UpgradeStoreLoader(upgradeInfo.Height, &storeUpgrades))
	}

	if loadLatest {
		if err := app.LoadLatestVersion(); err != nil {
			tmos.Exit(err.Error())
		}

		ctx := app.BaseApp.NewUncachedContext(true, tmproto.Header{})
		if err := app.wasmKeeper.InitializePinnedCodes(ctx); err != nil {
			tmos.Exit(fmt.Sprintf("failed initialize pinned codes %s", err))
		}
	}

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
		ibcfee.NewAppModule(app.ibcFeeKeeper),
		ibctransfer.NewAppModule(app.ibcTransferKeeper),
		params.NewAppModule(app.paramsKeeper),
		mint.NewAppModule(app.cdc, app.mintKeeper, app.accountKeeper),
		slashing.NewAppModule(app.cdc, app.slashingKeeper, app.accountKeeper, app.bankKeeper, app.stakingKeeper),
		staking.NewAppModule(app.cdc, app.stakingKeeper, app.accountKeeper, app.bankKeeper),
		custommint.NewAppModule(cdc, app.customMintKeeper),
		swap.NewAppModule(app.cdc, app.swapKeeper),
		vpn.NewAppModule(app.cdc, app.accountKeeper, app.bankKeeper, app.vpnKeeper),
		wasm.NewAppModule(app.cdc, &app.wasmKeeper, app.stakingKeeper, app.accountKeeper, app.bankKeeper),
	)
	app.simulationManager.RegisterStoreDecoders()

	return app
}

func (app *App) Name() string { return app.BaseApp.Name() }

func (app *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.moduleManager.BeginBlock(ctx, req)
}

func (app *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.moduleManager.EndBlock(ctx, req)
}

func (app *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var state map[string]json.RawMessage
	if err := tmjson.Unmarshal(req.AppStateBytes, &state); err != nil {
		panic(err)
	}

	app.upgradeKeeper.SetModuleVersionMap(ctx, app.moduleManager.GetVersionMap())
	return app.moduleManager.InitGenesis(ctx, app.cdc, state)
}

func (app *App) LegacyAmino() *codec.LegacyAmino {
	return app.amino
}

func (app *App) AppCodec() codec.Codec {
	return app.cdc
}

func (app *App) InterfaceRegistry() codectypes.InterfaceRegistry {
	return app.interfaceRegistry
}

func (app *App) RegisterTxService(ctx client.Context) {
	authtx.RegisterTxService(app.BaseApp.GRPCQueryRouter(), ctx, app.BaseApp.Simulate, app.interfaceRegistry)
}

func (app *App) RegisterTendermintService(ctx client.Context) {
	tmservice.RegisterTendermintService(app.BaseApp.GRPCQueryRouter(), ctx, app.interfaceRegistry)
}

func (app *App) RegisterAPIRoutes(server *api.Server, _ serverconfig.APIConfig) {
	ctx := server.ClientCtx
	rpc.RegisterRoutes(ctx, server.Router)
	authrest.RegisterTxRoutes(ctx, server.Router)
	authtx.RegisterGRPCGatewayRoutes(ctx, server.GRPCGatewayRouter)
	tmservice.RegisterGRPCGatewayRoutes(ctx, server.GRPCGatewayRouter)

	ModuleBasics.RegisterRESTRoutes(ctx, server.Router)
	ModuleBasics.RegisterGRPCGatewayRoutes(ctx, server.GRPCGatewayRouter)
}

func (app *App) LoadHeight(height int64) error {
	return app.LoadVersion(height)
}

func (app *App) ModuleAccountPermissions() map[string][]string {
	return map[string][]string{
		authtypes.FeeCollectorName:     nil,
		distributiontypes.ModuleName:   nil,
		govtypes.ModuleName:            {authtypes.Burner},
		minttypes.ModuleName:           {authtypes.Minter},
		stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
		stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
		ibcicatypes.ModuleName:         nil,
		ibcfeetypes.ModuleName:         nil,
		ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
		customminttypes.ModuleName:     nil,
		deposittypes.ModuleName:        nil,
		swaptypes.ModuleName:           {authtypes.Minter},
		wasmtypes.ModuleName:           {authtypes.Burner},
	}
}

func (app *App) ModuleAccountAddrs() map[string]bool {
	accounts := make(map[string]bool)
	for name := range app.ModuleAccountPermissions() {
		accounts[authtypes.NewModuleAddress(name).String()] = true
	}

	return accounts
}

func (app *App) SimulationManager() *module.SimulationManager {
	return app.simulationManager
}

func (app *App) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := app.paramsKeeper.GetSubspace(moduleName)
	return subspace
}
