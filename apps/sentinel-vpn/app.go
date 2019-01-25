package sentinel_vpn

import (
	"encoding/json"
	"io"
	"os"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/params"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/common"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmTypes "github.com/tendermint/tendermint/types"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/hub"
	"github.com/ironman0x7b2/sentinel-sdk/x/ibc"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

const appName = "Sentinel VPN"

var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.sentinel-vpn-cli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.sentinel-vpnd")
)

type SentinelVPN struct {
	*baseapp.BaseApp
	cdc *codec.Codec

	keyMain          *csdkTypes.KVStoreKey
	keyAccount       *csdkTypes.KVStoreKey
	keyFeeCollection *csdkTypes.KVStoreKey
	keyParams        *csdkTypes.KVStoreKey
	keyIBC           *csdkTypes.KVStoreKey
	keyVPN           *csdkTypes.KVStoreKey
	keySession       *csdkTypes.KVStoreKey

	tkeyParams *csdkTypes.TransientStoreKey

	accountKeeper       auth.AccountKeeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	bankKeeper          bank.Keeper
	paramsKeeper        params.Keeper
	ibcKeeper           ibc.Keeper
	vpnKeeper           vpn.Keeper
}

func NewSentinelVPN(logger log.Logger, db tmDB.DB, traceStore io.Writer, loadLatest bool, baseAppOptions ...func(*baseapp.BaseApp)) *SentinelVPN {
	cdc := MakeCodec()

	bApp := baseapp.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)

	var app = &SentinelVPN{
		BaseApp:          bApp,
		keyMain:          csdkTypes.NewKVStoreKey(baseapp.MainStoreKey),
		keyAccount:       csdkTypes.NewKVStoreKey(auth.StoreKey),
		keyFeeCollection: csdkTypes.NewKVStoreKey(auth.FeeStoreKey),
		keyParams:        csdkTypes.NewKVStoreKey(params.StoreKey),
		tkeyParams:       csdkTypes.NewTransientStoreKey(params.TStoreKey),
		keyIBC:           csdkTypes.NewKVStoreKey(sdkTypes.KeyIBC),
		keyVPN:           csdkTypes.NewKVStoreKey(sdkTypes.KeyVPN),
		keySession:       csdkTypes.NewKVStoreKey(sdkTypes.KeySession),
	}

	app.paramsKeeper = params.NewKeeper(app.cdc, app.keyParams, app.tkeyParams)
	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		app.keyAccount,
		app.paramsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount,
	)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper)
	app.ibcKeeper = ibc.NewKeeper(app.keyIBC, app.cdc)
	app.vpnKeeper = vpn.NewKeeper(app.cdc, app.keyVPN, app.keySession)

	app.Router().
		AddRoute(bank.RouterKey, bank.NewHandler(app.bankKeeper)).
		AddRoute(sdkTypes.KeyVPN, vpn.NewHandler(app.vpnKeeper, app.ibcKeeper)).
		AddRoute(sdkTypes.KeyIBC, vpn.NewIBCVPNHandler(app.ibcKeeper, app.vpnKeeper))

	app.MountStores(app.keyMain, app.keyAccount, app.keyFeeCollection, app.keyParams, app.keyIBC, app.keyVPN, app.keySession)
	app.SetInitChainer(app.initChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper))
	app.MountStoresTransient(app.tkeyParams)
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		err := app.LoadLatestVersion(app.keyMain)
		if err != nil {
			common.Exit(err.Error())
		}
	}

	app.Seal()

	return app
}

func MakeCodec() *codec.Codec {
	cdc := codec.New()

	codec.RegisterCrypto(cdc)
	csdkTypes.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	sdkTypes.RegisterCodec(cdc)
	ibc.RegisterCodec(cdc)
	hub.RegisterCodec(cdc)
	vpn.RegisterCodec(cdc)

	cdc.Seal()

	return cdc
}

func (app *SentinelVPN) BeginBlocker(_ csdkTypes.Context, _ abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	return abciTypes.ResponseBeginBlock{}
}

func (app *SentinelVPN) EndBlocker(_ csdkTypes.Context, _ abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	return abciTypes.ResponseEndBlock{}
}

func (app *SentinelVPN) initChainer(ctx csdkTypes.Context, req abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
	stateJSON := req.AppStateBytes

	genesisState := new(sdkTypes.GenesisState)
	err := app.cdc.UnmarshalJSON(stateJSON, genesisState)
	if err != nil {
		panic(err)
	}

	for _, gacc := range genesisState.Accounts {
		acc, err := gacc.ToAppAccount()
		if err != nil {
			panic(err)
		}

		acc.AccountNumber = app.accountKeeper.GetNextAccountNumber(ctx)
		app.accountKeeper.SetAccount(ctx, acc)
	}

	return abciTypes.ResponseInitChain{}
}

func (app *SentinelVPN) ExportAppStateAndValidators() (appState json.RawMessage, validators []tmTypes.GenesisValidator, err error) {
	ctx := app.NewContext(true, abciTypes.Header{})
	var accounts []*sdkTypes.GenesisAccount

	appendAccountsFn := func(acc auth.Account) bool {
		account := &sdkTypes.GenesisAccount{
			Address: acc.GetAddress(),
			Coins:   acc.GetCoins(),
		}

		accounts = append(accounts, account)
		return false
	}

	app.accountKeeper.IterateAccounts(ctx, appendAccountsFn)

	genState := sdkTypes.GenesisState{Accounts: accounts}
	appState, err = codec.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}

	return appState, validators, err
}
