package app

import (
	"encoding/json"
	"os"

	"github.com/cosmos/cosmos-sdk/baseapp"
	ccsdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/common"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/ibc"
	"github.com/ironman0x7b2/sentinel-sdk/types"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	tmDb "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmTypes "github.com/tendermint/tendermint/types"
)

const appName = "Sentinel Hub"

var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.sentinel-sdk-cli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.sentinel-sdkd")
)

type SentinelHub struct {
	*baseapp.BaseApp
	cdc *codec.Codec

	keyMain    *ccsdkTypes.KVStoreKey
	keyAccount *ccsdkTypes.KVStoreKey
	keyIBC     *ccsdkTypes.KVStoreKey

	accountKeeper       auth.AccountKeeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	bankKeeper          bank.Keeper
	ibcMapper           ibc.Mapper
	vpnKeeper           vpn.Keeper
}

func NewSentinelHub(logger log.Logger, db tmDb.DB, baseAppOptions ...func(*baseapp.BaseApp)) *SentinelHub {
	cdc := MakeCodec()

	var app = &SentinelHub{
		cdc:        cdc,
		BaseApp:    baseapp.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...),
		keyMain:    ccsdkTypes.NewKVStoreKey("main"),
		keyAccount: ccsdkTypes.NewKVStoreKey("acc"),
		keyIBC:     ccsdkTypes.NewKVStoreKey("ibc"),
	}

	app.accountKeeper = auth.NewAccountKeeper(
		cdc,
		app.keyAccount,
		func() auth.Account {
			return &types.AppAccount{}
		},
	)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper)
	app.ibcMapper = ibc.NewMapper(app.cdc, app.keyIBC, app.RegisterCodespace(ibc.DefaultCodespace))
	app.vpnKeeper = vpn.NewKeeper(app.keyVpn, app.keyIBC)

	app.Router().
		AddRoute("bank", bank.NewHandler(app.bankKeeper)).
		AddRoute("ibc", ibc.NewHandler(app.ibcMapper, app.bankKeeper)).
		AddRoute("vpn", vpn.NewHandler(app.vpnKeeper, app.ibcMapper))

	app.SetInitChainer(app.initChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetEndBlocker(app.EndBlocker)
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper))

	app.MountStoresIAVL(app.keyMain, app.keyAccount, app.keyIBC, app.keyVpn)
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
	ccsdkTypes.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	ibc.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	vpn.RegisterCodec(cdc)

	cdc.RegisterConcrete(&types.AppAccount{}, "sentinel-sdk/Account", nil)

	cdc.Seal()

	return cdc
}

func (app *SentinelHub) BeginBlocker(_ ccsdkTypes.Context, _ abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	return abciTypes.ResponseBeginBlock{}
}

func (app *SentinelHub) EndBlocker(_ ccsdkTypes.Context, _ abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	return abciTypes.ResponseEndBlock{}
}

func (app *SentinelHub) initChainer(ctx ccsdkTypes.Context, req abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
	stateJSON := req.AppStateBytes

	genesisState := new(types.GenesisState)
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

func (app *SentinelHub) ExportAppStateAndValidators() (appState json.RawMessage, validators []tmTypes.GenesisValidator, err error) {
	ctx := app.NewContext(true, abciTypes.Header{})
	accounts := []*types.GenesisAccount{}

	appendAccountsFn := func(acc auth.Account) bool {
		account := &types.GenesisAccount{
			Address: acc.GetAddress(),
			Coins:   acc.GetCoins(),
		}

		accounts = append(accounts, account)
		return false
	}

	app.accountKeeper.IterateAccounts(ctx, appendAccountsFn)

	genState := types.GenesisState{Accounts: accounts}
	appState, err = codec.MarshalJSONIndent(app.cdc, genState)
	if err != nil {
		return nil, nil, err
	}

	return appState, validators, err
}