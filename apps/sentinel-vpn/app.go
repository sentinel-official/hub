package sentinel_vpn

import (
	"encoding/json"
	"os"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/common"
	tmDb "github.com/tendermint/tendermint/libs/db"
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

	keyMain    *csdkTypes.KVStoreKey
	keyAccount *csdkTypes.KVStoreKey
	keyIBC     *csdkTypes.KVStoreKey
	keyVPN     *csdkTypes.KVStoreKey
	keySession *csdkTypes.KVStoreKey

	accountKeeper       auth.AccountKeeper
	bankKeeper          bank.Keeper
	ibcKeeper           ibc.Keeper
	vpnKeeper           vpn.Keeper
	feeCollectionKeeper auth.FeeCollectionKeeper
}

func NewSentinelVPN(logger log.Logger, db tmDb.DB, baseAppOptions ...func(*baseapp.BaseApp)) *SentinelVPN {
	cdc := MakeCodec()

	var app = &SentinelVPN{
		cdc:        cdc,
		BaseApp:    baseapp.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...),
		keyMain:    csdkTypes.NewKVStoreKey(sdkTypes.KeyMain),
		keyAccount: csdkTypes.NewKVStoreKey(sdkTypes.KeyAccount),
		keyIBC:     csdkTypes.NewKVStoreKey(sdkTypes.KeyIBC),
		keyVPN:     csdkTypes.NewKVStoreKey(sdkTypes.KeyVPN),
		keySession: csdkTypes.NewKVStoreKey(sdkTypes.KeySession),
	}

	app.accountKeeper = auth.NewAccountKeeper(
		cdc,
		app.keyAccount,
		func() auth.Account {
			return &sdkTypes.AppAccount{}
		},
	)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper)
	app.ibcKeeper = ibc.NewKeeper(app.keyIBC, app.cdc)
	app.vpnKeeper = vpn.NewKeeper(app.cdc, app.keyVPN, app.keySession)

	app.Router().
		AddRoute(sdkTypes.KeyBank, bank.NewHandler(app.bankKeeper)).
		AddRoute(sdkTypes.KeyVPN, vpn.NewHandler(app.vpnKeeper, app.ibcKeeper)).
		AddRoute(sdkTypes.KeyIBC, vpn.NewIBCVPNHandler(app.ibcKeeper, app.vpnKeeper))

	app.SetInitChainer(app.initChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper))

	app.MountStores(app.keyMain, app.keyAccount, app.keyIBC, app.keyVPN, app.keySession)
	err := app.LoadLatestVersion(app.keyMain)
	if err != nil {
		common.Exit(err.Error())
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
		// TODO: https://github.com/cosmos/cosmos-sdk/issues/468
		panic(err)
	}

	for _, gacc := range genesisState.Accounts {
		acc, err := gacc.ToAppAccount()
		if err != nil {
			// TODO: https://github.com/cosmos/cosmos-sdk/issues/468
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
