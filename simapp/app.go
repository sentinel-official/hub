package simapp

import (
	"encoding/json"
	"io"
	"os"
	
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
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
	"github.com/sentinel-official/hub/version"
	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/vpn"
)

const (
	appName = "SimApp"
)

var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.simapp")
	DefaultNodeHome = os.ExpandEnv("$HOME/.simapp")
	
	ModuleBasics = module.NewBasicManager(
		genaccounts.AppModuleBasic{},
		genutil.AppModuleBasic{},
		auth.AppModuleBasic{},
		bank.AppModuleBasic{},
		staking.AppModuleBasic{},
		mint.AppModuleBasic{},
		distribution.AppModuleBasic{},
		gov.NewAppModuleBasic(client.ProposalHandler, distribution.ProposalHandler),
		params.AppModuleBasic{},
		crisis.AppModuleBasic{},
		slashing.AppModuleBasic{},
		supply.AppModuleBasic{},
		deposit.AppModuleBasic{},
		vpn.AppModuleBasic{},
	)
	
	moduleAccountPermissions = map[string][]string{
		auth.FeeCollectorName:     nil,
		distribution.ModuleName:   nil,
		mint.ModuleName:           {supply.Minter},
		staking.BondedPoolName:    {supply.Burner, supply.Staking},
		staking.NotBondedPoolName: {supply.Burner, supply.Staking},
		gov.ModuleName:            {supply.Burner},
		deposit.ModuleName:        nil,
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

func SetBech32AddressPrefixes(config *sdk.Config) {
	config.SetBech32PrefixForAccount(types.Bech32PrefixAccAddr, types.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(types.Bech32PrefixValAddr, types.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(types.Bech32PrefixConsAddr, types.Bech32PrefixConsPub)
}

type SimApp struct {
	*baseapp.BaseApp
	cdc *codec.Codec
	
	invCheckPeriod uint
	
	keys          map[string]*sdk.KVStoreKey
	transientKeys map[string]*sdk.TransientStoreKey
	
	accountKeeper      auth.AccountKeeper
	bankKeeper         bank.Keeper
	supplyKeeper       supply.Keeper
	stakingKeeper      staking.Keeper
	slashingKeeper     slashing.Keeper
	mintKeeper         mint.Keeper
	distributionKeeper distribution.Keeper
	govKeeper          gov.Keeper
	crisisKeeper       crisis.Keeper
	paramsKeeper       params.Keeper
	depositKeeper      deposit.Keeper
	vpnKeeper          vpn.Keeper
	
	mm *module.Manager
}

// nolint:funlen
func NewSimApp(logger log.Logger, db db.DB,
	traceStore io.Writer, loadLatest bool, invCheckPeriod uint,
	baseAppOptions ...func(*baseapp.BaseApp)) *SimApp {
	cdc := MakeCodec()
	
	bApp := baseapp.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)
	bApp.SetAppVersion(version.Version)
	
	keys := sdk.NewKVStoreKeys(
		baseapp.MainStoreKey, auth.StoreKey, staking.StoreKey,
		supply.StoreKey, mint.StoreKey, distribution.StoreKey, slashing.StoreKey,
		gov.StoreKey, params.StoreKey, deposit.StoreKey,
		vpn.StoreKeyNode, vpn.StoreKeySubscription, vpn.StoreKeySession,
	)
	
	transientKeys := sdk.NewTransientStoreKeys(staking.TStoreKey, params.TStoreKey)
	
	var app = &SimApp{
		BaseApp:        bApp,
		cdc:            cdc,
		invCheckPeriod: invCheckPeriod,
		keys:           keys,
		transientKeys:  transientKeys,
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
	govSubspace := app.paramsKeeper.Subspace(gov.DefaultParamspace)
	crisisSubspace := app.paramsKeeper.Subspace(crisis.DefaultParamspace)
	
	app.accountKeeper = auth.NewAccountKeeper(app.cdc,
		keys[auth.StoreKey],
		authSubspace,
		auth.ProtoBaseAccount)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper,
		bankSubspace,
		bank.DefaultCodespace,
		app.ModuleAccountAddrs())
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
		app.ModuleAccountAddrs())
	app.slashingKeeper = slashing.NewKeeper(app.cdc,
		keys[slashing.StoreKey],
		&stakingKeeper,
		slashingSubspace,
		slashing.DefaultCodespace)
	app.crisisKeeper = crisis.NewKeeper(crisisSubspace,
		invCheckPeriod,
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
	
	app.depositKeeper = deposit.NewKeeper(app.cdc,
		keys[deposit.StoreKey],
		app.supplyKeeper)
	app.vpnKeeper = vpn.NewKeeper(app.cdc,
		keys[vpn.StoreKeyNode],
		keys[vpn.StoreKeySubscription],
		keys[vpn.StoreKeySession],
		keys[vpn.StoreKeyResolver],
		app.paramsKeeper.Subspace(vpn.DefaultParamspace),
		app.depositKeeper)
	
	app.mm = module.NewManager(
		genaccounts.NewAppModule(app.accountKeeper),
		genutil.NewAppModule(app.accountKeeper, app.stakingKeeper, app.BaseApp.DeliverTx),
		auth.NewAppModule(app.accountKeeper),
		bank.NewAppModule(app.bankKeeper, app.accountKeeper),
		crisis.NewAppModule(&app.crisisKeeper),
		supply.NewAppModule(app.supplyKeeper, app.accountKeeper),
		distribution.NewAppModule(app.distributionKeeper, app.supplyKeeper),
		gov.NewAppModule(app.govKeeper, app.supplyKeeper),
		mint.NewAppModule(app.mintKeeper),
		slashing.NewAppModule(app.slashingKeeper, app.stakingKeeper),
		staking.NewAppModule(app.stakingKeeper, app.distributionKeeper, app.accountKeeper, app.supplyKeeper),
		deposit.NewAppModule(app.depositKeeper),
		vpn.NewAppModule(app.vpnKeeper),
	)
	
	app.mm.SetOrderBeginBlockers(mint.ModuleName, distribution.ModuleName, slashing.ModuleName)
	app.mm.SetOrderEndBlockers(crisis.ModuleName, gov.ModuleName, staking.ModuleName, vpn.ModuleName)
	app.mm.SetOrderInitGenesis(
		genaccounts.ModuleName, distribution.ModuleName, staking.ModuleName,
		auth.ModuleName, bank.ModuleName, slashing.ModuleName, gov.ModuleName,
		mint.ModuleName, supply.ModuleName, crisis.ModuleName, genutil.ModuleName,
		deposit.ModuleName, vpn.ModuleName,
	)
	
	app.mm.RegisterInvariants(&app.crisisKeeper)
	app.mm.RegisterRoutes(app.Router(), app.QueryRouter())
	app.MountKVStores(keys)
	app.MountTransientStores(transientKeys)
	
	app.SetInitChainer(app.InitChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(
		auth.NewAnteHandler(app.accountKeeper, app.supplyKeeper, auth.DefaultSigVerificationGasConsumer))
	app.SetEndBlocker(app.EndBlocker)
	
	if loadLatest {
		if err := app.LoadLatestVersion(app.keys[baseapp.MainStoreKey]); err != nil {
			common.Exit(err.Error())
		}
	}
	
	return app
}

func (app *SimApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {
	return app.mm.BeginBlock(ctx, req)
}

func (app *SimApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
	return app.mm.EndBlock(ctx, req)
}

func (app *SimApp) InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	var state map[string]json.RawMessage
	app.cdc.MustUnmarshalJSON(req.AppStateBytes, &state)
	
	return app.mm.InitGenesis(ctx, state)
}

func (app *SimApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keys[baseapp.MainStoreKey])
}

func (app *SimApp) ModuleAccountAddrs() map[string]bool {
	moduleAccounts := make(map[string]bool)
	for acc := range moduleAccountPermissions {
		moduleAccounts[supply.NewModuleAddress(acc).String()] = true
	}
	
	return moduleAccounts
}
