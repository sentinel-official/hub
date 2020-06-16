package app

import (
	"encoding/json"
	"io"
	"os"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/genaccounts"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/params/client"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/supply"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
	db "github.com/tendermint/tm-db"

	"github.com/sentinel-official/hub/types"
	"github.com/sentinel-official/hub/x/dvpn"
)

const (
	appName = "Sentinel Hub"
)

var (
	DefaultNodeHome = os.ExpandEnv("$HOME/.sentinel-hub-daemon")
	DefaultCLIHome  = os.ExpandEnv("$HOME/.sentinel-hub-cli")

	ModuleBasics = module.NewBasicManager(
		genaccounts.AppModuleBasic{},
		genutil.AppModuleBasic{},
		params.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		supply.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distribution.AppModuleBasic{},
		slashing.AppModuleBasic{},
		crisis.AppModuleBasic{},
		gov.NewAppModuleBasic(client.ProposalHandler, distribution.ProposalHandler),
		dvpn.AppModuleBasic{},
	)

	moduleAccountPermissions = map[string][]string{
		auth.FeeCollectorName:     nil,
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
		mint.ModuleName:           {supply.Minter},
		distribution.ModuleName:   nil,
		gov.ModuleName:            {supply.Burner},
	}
)

func MakeCodec() *codec.Codec {
	var cdc = codec.New()

	sdk.RegisterCodec(cdc)
	types.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
	ModuleBasics.RegisterCodec(cdc)

	return cdc
}

type App struct {
	*baseapp.BaseApp

	cdc           *codec.Codec
	manager       *module.Manager
	keys          map[string]*sdk.KVStoreKey
	transientKeys map[string]*sdk.TransientStoreKey

	paramsKeeper       params.Keeper
	accountKeeper      auth.AccountKeeper
	bankKeeper         bank.Keeper
	supplyKeeper       supply.Keeper
	stakingKeeper      staking.Keeper
	mintKeeper         mint.Keeper
	distributionKeeper distribution.Keeper
	slashingKeeper     slashing.Keeper
	crisisKeeper       crisis.Keeper
	govKeeper          gov.Keeper
	dVPNKeeper         dvpn.Keeper
}

func NewApp(logger log.Logger, db db.DB, tracer io.Writer, latest bool, invarCheckPeriod uint,
	options ...func(*baseapp.BaseApp)) *App {
	cdc := MakeCodec()

	baseApp := baseapp.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), options...)
	baseApp.SetCommitMultiStoreTracer(tracer)
	baseApp.SetAppVersion(version.Version)

	keys := sdk.NewKVStoreKeys(baseapp.MainStoreKey,
		params.StoreKey, auth.StoreKey, supply.StoreKey, staking.StoreKey,
		mint.StoreKey, distribution.StoreKey, slashing.StoreKey, gov.StoreKey,
		dvpn.StoreKey,
	)
	transientKeys := sdk.NewTransientStoreKeys(params.TStoreKey, staking.TStoreKey)

	var app = &App{
		BaseApp:       baseApp,
		cdc:           cdc,
		keys:          keys,
		transientKeys: transientKeys,
	}

	app.paramsKeeper = params.NewKeeper(app.cdc,
		keys[params.StoreKey],
		transientKeys[params.TStoreKey],
		params.DefaultCodespace)

	authSubspace := app.paramsKeeper.Subspace(auth.DefaultParamspace)
	bankSubspace := app.paramsKeeper.Subspace(bank.DefaultParamspace)
	stakingSubspace := app.paramsKeeper.Subspace(staking.DefaultParamspace)
	mintSubspace := app.paramsKeeper.Subspace(mint.DefaultParamspace)
	distributionSubspace := app.paramsKeeper.Subspace(distribution.DefaultParamspace)
	slashingSubspace := app.paramsKeeper.Subspace(slashing.DefaultParamspace)
	crisisSubspace := app.paramsKeeper.Subspace(crisis.DefaultParamspace)
	govSubspace := app.paramsKeeper.Subspace(gov.DefaultParamspace)

	app.accountKeeper = auth.NewAccountKeeper(app.cdc,
		keys[auth.StoreKey],
		authSubspace,
		auth.ProtoBaseAccount)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper,
		bankSubspace,
		bank.DefaultCodespace,
		app.ModuleAccountAddresses())
	app.supplyKeeper = supply.NewKeeper(app.cdc,
		keys[supply.StoreKey],
		app.accountKeeper,
		app.bankKeeper,
		moduleAccountPermissions)
	stakingKeeper := staking.NewKeeper(app.cdc,
		keys[staking.StoreKey],
		transientKeys[staking.TStoreKey],
		app.supplyKeeper,
		stakingSubspace,
		staking.DefaultCodespace)
	app.mintKeeper = mint.NewKeeper(app.cdc,
		keys[mint.StoreKey],
		mintSubspace,
		&stakingKeeper,
		app.supplyKeeper,
		auth.FeeCollectorName)
	app.distributionKeeper = distribution.NewKeeper(app.cdc,
		keys[distribution.StoreKey],
		distributionSubspace,
		&stakingKeeper,
		app.supplyKeeper,
		distribution.DefaultCodespace,
		auth.FeeCollectorName,
		app.ModuleAccountAddresses())
	app.slashingKeeper = slashing.NewKeeper(app.cdc,
		keys[slashing.StoreKey],
		&stakingKeeper,
		slashingSubspace,
		slashing.DefaultCodespace)
	app.crisisKeeper = crisis.NewKeeper(crisisSubspace,
		invarCheckPeriod,
		app.supplyKeeper,
		auth.FeeCollectorName)

	govRouter := gov.NewRouter()
	govRouter.AddRoute(gov.RouterKey, gov.ProposalHandler).
		AddRoute(params.RouterKey, params.NewParamChangeProposalHandler(app.paramsKeeper)).
		AddRoute(distribution.RouterKey, distribution.NewCommunityPoolSpendProposalHandler(app.distributionKeeper))

	app.govKeeper = gov.NewKeeper(app.cdc,
		keys[gov.StoreKey],
		app.paramsKeeper,
		govSubspace,
		app.supplyKeeper,
		&stakingKeeper,
		gov.DefaultCodespace,
		govRouter)
	app.stakingKeeper = *stakingKeeper.SetHooks(
		staking.NewMultiStakingHooks(app.distributionKeeper.Hooks(), app.slashingKeeper.Hooks()))
	app.dVPNKeeper = dvpn.NewKeeper(app.cdc, keys[dvpn.StoreKey])

	app.manager = module.NewManager(
		genaccounts.NewAppModule(app.accountKeeper),
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		staking.NewAppModule(app.stakingKeeper, app.distributionKeeper, app.accountKeeper, app.supplyKeeper),
		mint.NewAppModule(app.mintKeeper),
		distribution.NewAppModule(app.distributionKeeper, app.supplyKeeper),
		slashing.NewAppModule(app.slashingKeeper, app.stakingKeeper),
		crisis.NewAppModule(&app.crisisKeeper),
		gov.NewAppModule(app.govKeeper, app.supplyKeeper),
		dvpn.NewAppModule(app.dVPNKeeper),
	)

	// NOTE: order is very important here
	app.manager.SetOrderBeginBlockers(mint.ModuleName, distribution.ModuleName, slashing.ModuleName)
	app.manager.SetOrderEndBlockers(crisis.ModuleName, gov.ModuleName, staking.ModuleName)
	app.manager.SetOrderInitGenesis(
		genaccounts.ModuleName, distribution.ModuleName, staking.ModuleName, auth.ModuleName,
		bank.ModuleName, slashing.ModuleName, gov.ModuleName, mint.ModuleName,
		supply.ModuleName, crisis.ModuleName, genutil.ModuleName, dvpn.ModuleName,
	)

	app.manager.RegisterInvariants(&app.crisisKeeper)
	app.manager.RegisterRoutes(app.Router(), app.QueryRouter())
	app.MountKVStores(keys)
	app.MountTransientStores(transientKeys)

	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.supplyKeeper, auth.DefaultSigVerificationGasConsumer))
	app.SetEndBlocker(app.EndBlocker)

	if latest {
		if err := app.LoadLatestVersion(app.keys[baseapp.MainStoreKey]); err != nil {
			common.Exit(err.Error())
		}
	}

	return app
}

func (a *App) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return a.manager.BeginBlock(ctx, req)
}

func (a *App) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return a.manager.EndBlock(ctx, req)
}

func (a *App) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var state map[string]json.RawMessage
	a.cdc.MustUnmarshalJSON(req.AppStateBytes, &state)

	return a.manager.InitGenesis(ctx, state)
}

func (a *App) LoadHeight(height int64) error {
	return a.LoadVersion(height, a.keys[baseapp.MainStoreKey])
}

func (a *App) ModuleAccountAddresses() map[string]bool {
	accounts := make(map[string]bool)
	for name := range moduleAccountPermissions {
		accounts[supply.NewModuleAddress(name).String()] = true
	}

	return accounts
}
