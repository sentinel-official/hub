package hub

import (
	"encoding/json"
	"io"
	"os"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramsclient "github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradeclient "github.com/cosmos/cosmos-sdk/x/upgrade/client"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"
	tmdb "github.com/tendermint/tm-db"

	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/vpn"
)

const (
	appName = "Sentinel Hub"
)

var (
	DefaultNodeHome = os.ExpandEnv("${HOME}/.sentinelhubd")
	DefaultCLIHome  = os.ExpandEnv("${HOME}/.sentinelhubcli")
	ModuleBasics    = module.NewBasicManager(
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		crisis.AppModuleBasic{},
		distribution.AppModuleBasic{},
		evidence.AppModuleBasic{},
		genutil.AppModuleBasic{},
		gov.NewAppModuleBasic(
			distribution.ProposalHandler,
			paramsclient.ProposalHandler,
			upgradeclient.ProposalHandler,
		),
		mint.AppModuleBasic{},
		params.AppModuleBasic{},
		slashing.AppModuleBasic{},
		staking.AppModuleBasic{},
		supply.AppModuleBasic{},
		upgrade.AppModuleBasic{},
		vpn.AppModuleBasic{},
	)
)

func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
	vesting.RegisterCodec(cdc)
	ModuleBasics.RegisterCodec(cdc)

	return cdc
}

var _ simapp.App = (*App)(nil)

type App struct {
	*baseapp.BaseApp

	cdc *codec.Codec
	mm  *module.Manager
	msm *module.SimulationManager

	keys      map[string]*sdk.KVStoreKey
	tkeys     map[string]*sdk.TransientStoreKey
	subspaces map[string]params.Subspace

	accountKeeper      auth.AccountKeeper
	bankKeeper         bank.Keeper
	crisisKeeper       crisis.Keeper
	distributionKeeper distribution.Keeper
	evidenceKeeper     evidence.Keeper
	govKeeper          gov.Keeper
	mintKeeper         mint.Keeper
	paramsKeeper       params.Keeper
	slashingKeeper     slashing.Keeper
	stakingKeeper      staking.Keeper
	supplyKeeper       supply.Keeper
	upgradeKeeper      upgrade.Keeper
	vpnKeeper          vpn.Keeper
}

func NewApp(
	logger log.Logger,
	db tmdb.DB,
	tracer io.Writer,
	loadLatest bool,
	skipUpgradeHeights map[int64]bool,
	invarCheckPeriod uint,
	options ...func(*baseapp.BaseApp),
) *App {
	var (
		cdc  = MakeCodec()
		keys = sdk.NewKVStoreKeys(baseapp.MainStoreKey,
			auth.StoreKey, distribution.StoreKey, evidence.StoreKey, gov.StoreKey,
			mint.StoreKey, params.StoreKey, slashing.StoreKey, staking.StoreKey,
			supply.StoreKey, upgrade.StoreKey, vpn.StoreKey,
		)
		tkeys = sdk.NewTransientStoreKeys(params.TStoreKey)
	)

	baseApp := baseapp.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), options...)
	baseApp.SetCommitMultiStoreTracer(tracer)
	baseApp.SetAppVersion(version.Version)

	app := &App{
		BaseApp:   baseApp,
		cdc:       cdc,
		keys:      keys,
		tkeys:     tkeys,
		subspaces: make(map[string]params.Subspace),
	}

	app.paramsKeeper = params.NewKeeper(app.cdc,
		keys[params.StoreKey],
		tkeys[params.TStoreKey])

	app.subspaces[auth.ModuleName] = app.paramsKeeper.Subspace(auth.DefaultParamspace)
	app.subspaces[bank.ModuleName] = app.paramsKeeper.Subspace(bank.DefaultParamspace)
	app.subspaces[crisis.ModuleName] = app.paramsKeeper.Subspace(crisis.DefaultParamspace)
	app.subspaces[distribution.ModuleName] = app.paramsKeeper.Subspace(distribution.DefaultParamspace)
	app.subspaces[evidence.ModuleName] = app.paramsKeeper.Subspace(evidence.DefaultParamspace)
	app.subspaces[gov.ModuleName] = app.paramsKeeper.Subspace(gov.DefaultParamspace).WithKeyTable(gov.ParamKeyTable())
	app.subspaces[mint.ModuleName] = app.paramsKeeper.Subspace(mint.DefaultParamspace)
	app.subspaces[slashing.ModuleName] = app.paramsKeeper.Subspace(slashing.DefaultParamspace)
	app.subspaces[staking.ModuleName] = app.paramsKeeper.Subspace(staking.DefaultParamspace)

	app.accountKeeper = auth.NewAccountKeeper(app.cdc,
		keys[auth.StoreKey],
		app.subspaces[auth.ModuleName],
		auth.ProtoBaseAccount)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper,
		app.subspaces[bank.ModuleName],
		app.ModuleAccountAddrsBlackList())
	app.supplyKeeper = supply.NewKeeper(app.cdc,
		keys[supply.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		app.ModuleAccountsPermissions())
	stakingKeeper := staking.NewKeeper(app.cdc,
		keys[staking.StoreKey],
		app.supplyKeeper,
		app.subspaces[staking.ModuleName])
	app.mintKeeper = mint.NewKeeper(app.cdc,
		keys[mint.StoreKey],
		app.subspaces[mint.ModuleName],
		&stakingKeeper,
		app.supplyKeeper,
		auth.FeeCollectorName)
	app.distributionKeeper = distribution.NewKeeper(app.cdc,
		keys[distribution.StoreKey],
		app.subspaces[distribution.ModuleName],
		&stakingKeeper,
		app.supplyKeeper,
		auth.FeeCollectorName,
		app.ModuleAccountAddrs())
	app.slashingKeeper = slashing.NewKeeper(app.cdc,
		keys[slashing.StoreKey],
		&stakingKeeper,
		app.subspaces[slashing.ModuleName])
	app.crisisKeeper = crisis.NewKeeper(app.subspaces[crisis.ModuleName],
		invarCheckPeriod,
		app.supplyKeeper,
		auth.FeeCollectorName)
	app.upgradeKeeper = upgrade.NewKeeper(skipUpgradeHeights,
		keys[upgrade.StoreKey],
		app.cdc)

	app.evidenceKeeper = *evidence.NewKeeper(app.cdc,
		keys[evidence.StoreKey],
		app.subspaces[evidence.ModuleName],
		&stakingKeeper,
		app.slashingKeeper,
	)

	evidenceRouter := evidence.NewRouter()
	app.evidenceKeeper.SetRouter(evidenceRouter)

	govRouter := gov.NewRouter()
	govRouter.AddRoute(gov.RouterKey, gov.ProposalHandler).
		AddRoute(params.RouterKey, params.NewParamChangeProposalHandler(app.paramsKeeper)).
		AddRoute(distribution.RouterKey, distribution.NewCommunityPoolSpendProposalHandler(app.distributionKeeper)).
		AddRoute(upgrade.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(app.upgradeKeeper))

	app.govKeeper = gov.NewKeeper(app.cdc,
		keys[gov.StoreKey],
		app.subspaces[gov.ModuleName],
		app.supplyKeeper,
		&stakingKeeper,
		govRouter)
	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(
			app.distributionKeeper.Hooks(),
			app.slashingKeeper.Hooks(),
		),
	)
	app.vpnKeeper = vpn.NewKeeper(app.cdc,
		keys[vpn.StoreKey],
		app.paramsKeeper,
		app.bankKeeper,
		app.supplyKeeper)

	app.mm = module.NewManager(
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		crisis.NewAppModule(&app.crisisKeeper),
		distribution.NewAppModule(app.distributionKeeper, app.accountKeeper, app.supplyKeeper, app.stakingKeeper),
		evidence.NewAppModule(app.evidenceKeeper),
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		gov.NewAppModule(app.govKeeper, app.accountKeeper, app.supplyKeeper),
		mint.NewAppModule(app.mintKeeper),
		slashing.NewAppModule(app.slashingKeeper, app.accountKeeper, app.stakingKeeper),
		staking.NewAppModule(app.stakingKeeper, app.accountKeeper, app.supplyKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		upgrade.NewAppModule(app.upgradeKeeper),
		vpn.NewAppModule(app.accountKeeper, app.vpnKeeper),
	)

	// NOTE: order is very important here
	app.mm.SetOrderBeginBlockers(
		upgrade.ModuleName, mint.ModuleName, distribution.ModuleName,
		slashing.ModuleName, evidence.ModuleName)
	app.mm.SetOrderEndBlockers(
		crisis.ModuleName, gov.ModuleName, staking.ModuleName, vpn.ModuleName)
	app.mm.SetOrderInitGenesis(
		auth.ModuleName, distribution.ModuleName, staking.ModuleName, bank.ModuleName,
		slashing.ModuleName, gov.ModuleName, mint.ModuleName, supply.ModuleName,
		crisis.ModuleName, genutil.ModuleName, evidence.ModuleName, vpn.ModuleName,
	)

	app.mm.RegisterInvariants(&app.crisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())

	app.msm = module.NewSimulationManager(
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		distribution.NewAppModule(app.distributionKeeper, app.accountKeeper, app.supplyKeeper, app.stakingKeeper),
		gov.NewAppModule(app.govKeeper, app.accountKeeper, app.supplyKeeper),
		mint.NewAppModule(app.mintKeeper),
		params.NewAppModule(),
		slashing.NewAppModule(app.slashingKeeper, app.accountKeeper, app.stakingKeeper),
		staking.NewAppModule(app.stakingKeeper, app.accountKeeper, app.supplyKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		vpn.NewAppModule(app.accountKeeper, app.vpnKeeper),
	)

	app.msm.RegisterStoreDecoders()
	app.MountKVStores(keys)
	app.MountTransientStores(tkeys)

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.supplyKeeper, auth.DefaultSigVerificationGasConsumer))
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(app.keys[baseapp.MainStoreKey]); err != nil {
			tmos.Exit(err.Error())
		}
	}

	return app
}

func (a *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return a.mm.BeginBlock(ctx, req)
}

func (a *App) Codec() *codec.Codec {
	return a.cdc
}

func (a *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return a.mm.EndBlock(ctx, req)
}

func (a *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var state map[string]json.RawMessage
	a.cdc.MustUnmarshalJSON(req.AppStateBytes, &state)

	return a.mm.InitGenesis(ctx, state)
}

func (a *App) LoadHeight(height int64) error {
	return a.LoadVersion(height, a.keys[baseapp.MainStoreKey])
}

func (a *App) ModuleAccountsPermissions() map[string][]string {
	return map[string][]string{
		auth.FeeCollectorName:     nil,
		distribution.ModuleName:   nil,
		gov.ModuleName:            {supply.Burner},
		mint.ModuleName:           {supply.Minter},
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
		deposit.ModuleName:        nil,
	}
}

func (a *App) ModuleAccountAddrsWhiteList() map[string]bool {
	accounts := make(map[string]bool)
	accounts[supply.NewModuleAddress(distribution.ModuleName).String()] = true

	return accounts
}

func (a *App) ModuleAccountAddrs() map[string]bool {
	accounts := make(map[string]bool)
	for name := range a.ModuleAccountsPermissions() {
		accounts[supply.NewModuleAddress(name).String()] = true
	}

	return accounts
}

func (a *App) ModuleAccountAddrsBlackList() map[string]bool {
	accounts := a.ModuleAccountAddrs()
	for name := range a.ModuleAccountAddrsWhiteList() {
		delete(accounts, supply.NewModuleAddress(name).String())
	}

	return accounts
}

func (a *App) SimulationManager() *module.SimulationManager {
	return a.msm
}
