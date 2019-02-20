package vpn

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/gov"
	"github.com/cosmos/cosmos-sdk/x/ibc"
	"github.com/cosmos/cosmos-sdk/x/mint"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	abciTypes "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/common"
	tmDB "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/ironman0x7b2/sentinel-sdk/x/vpn"
)

const (
	appName        = "Sentinel VPN"
	DefaultKeyPass = "1234567890"
)

var (
	DefaultCLIHome  = os.ExpandEnv("$HOME/.vpncli")
	DefaultNodeHome = os.ExpandEnv("$HOME/.vpnd")
)

type VPN struct {
	*baseapp.BaseApp
	cdc *codec.Codec

	keyMain *csdkTypes.KVStoreKey

	keyParams        *csdkTypes.KVStoreKey
	keyAccount       *csdkTypes.KVStoreKey
	keyFeeCollection *csdkTypes.KVStoreKey
	keyStaking       *csdkTypes.KVStoreKey
	keySlashing      *csdkTypes.KVStoreKey
	keyDistribution  *csdkTypes.KVStoreKey
	keyGov           *csdkTypes.KVStoreKey
	keyMint          *csdkTypes.KVStoreKey
	keyIBC           *csdkTypes.KVStoreKey

	keyVPNNode    *csdkTypes.KVStoreKey
	keyVPNSession *csdkTypes.KVStoreKey

	tkeyParams       *csdkTypes.TransientStoreKey
	tkeyStaking      *csdkTypes.TransientStoreKey
	tkeyDistribution *csdkTypes.TransientStoreKey

	paramsKeeper        params.Keeper
	accountKeeper       auth.AccountKeeper
	bankKeeper          bank.Keeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	stakingKeeper       staking.Keeper
	slashingKeeper      slashing.Keeper
	distributionKeeper  distribution.Keeper
	govKeeper           gov.Keeper
	mintKeeper          mint.Keeper
	ibcMapper           ibc.Mapper

	vpnKeeper vpn.Keeper
}

func NewVPN(logger log.Logger, db tmDB.DB, traceStore io.Writer, loadLatest bool, baseAppOptions ...func(*baseapp.BaseApp)) *VPN {
	cdc := MakeCodec()

	bApp := baseapp.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)

	var app = &VPN{
		BaseApp:          bApp,
		cdc:              cdc,
		keyParams:        csdkTypes.NewKVStoreKey(params.StoreKey),
		keyMain:          csdkTypes.NewKVStoreKey(baseapp.MainStoreKey),
		keyAccount:       csdkTypes.NewKVStoreKey(auth.StoreKey),
		keyFeeCollection: csdkTypes.NewKVStoreKey(auth.FeeStoreKey),
		keyStaking:       csdkTypes.NewKVStoreKey(staking.StoreKey),
		keySlashing:      csdkTypes.NewKVStoreKey(slashing.StoreKey),
		keyDistribution:  csdkTypes.NewKVStoreKey(distribution.StoreKey),
		keyGov:           csdkTypes.NewKVStoreKey(gov.StoreKey),
		keyMint:          csdkTypes.NewKVStoreKey(mint.StoreKey),
		keyIBC:           csdkTypes.NewKVStoreKey("ibc"),
		keyVPNNode:       csdkTypes.NewKVStoreKey(vpn.StoreKeyNode),
		keyVPNSession:    csdkTypes.NewKVStoreKey(vpn.StoreKeySession),
		tkeyParams:       csdkTypes.NewTransientStoreKey(params.TStoreKey),
		tkeyStaking:      csdkTypes.NewTransientStoreKey(staking.TStoreKey),
		tkeyDistribution: csdkTypes.NewTransientStoreKey(distribution.TStoreKey),
	}

	app.paramsKeeper = params.NewKeeper(app.cdc,
		app.keyParams,
		app.tkeyParams)
	app.accountKeeper = auth.NewAccountKeeper(app.cdc,
		app.keyAccount,
		app.paramsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount)
	app.bankKeeper = bank.NewBaseKeeper(app.accountKeeper,
		app.paramsKeeper.Subspace(bank.DefaultParamspace),
		bank.DefaultCodespace)
	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(app.cdc,
		app.keyFeeCollection)
	stakingKeeper := staking.NewKeeper(app.cdc,
		app.keyStaking,
		app.tkeyStaking,
		app.bankKeeper,
		app.paramsKeeper.Subspace(staking.DefaultParamspace),
		staking.DefaultCodespace)
	app.slashingKeeper = slashing.NewKeeper(app.cdc,
		app.keySlashing,
		&stakingKeeper,
		app.paramsKeeper.Subspace(slashing.DefaultParamspace),
		slashing.DefaultCodespace)
	app.distributionKeeper = distribution.NewKeeper(app.cdc,
		app.keyDistribution,
		app.paramsKeeper.Subspace(distribution.DefaultParamspace),
		app.bankKeeper,
		&stakingKeeper,
		app.feeCollectionKeeper,
		distribution.DefaultCodespace)
	app.govKeeper = gov.NewKeeper(app.cdc,
		app.keyGov,
		app.paramsKeeper,
		app.paramsKeeper.Subspace(gov.DefaultParamspace),
		app.bankKeeper,
		&stakingKeeper,
		gov.DefaultCodespace)
	app.mintKeeper = mint.NewKeeper(app.cdc,
		app.keyMint,
		app.paramsKeeper.Subspace(mint.DefaultParamspace),
		&stakingKeeper,
		app.feeCollectionKeeper)
	app.stakingKeeper = *stakingKeeper.SetHooks(NewStakingHooks(app.distributionKeeper.Hooks(), app.slashingKeeper.Hooks()))
	app.vpnKeeper = vpn.NewKeeper(app.cdc, app.keyVPNNode, app.keyVPNSession)

	app.Router().
		AddRoute(bank.RouterKey, bank.NewHandler(app.bankKeeper)).
		AddRoute(staking.RouterKey, staking.NewHandler(app.stakingKeeper)).
		AddRoute(slashing.RouterKey, slashing.NewHandler(app.slashingKeeper)).
		AddRoute(distribution.RouterKey, distribution.NewHandler(app.distributionKeeper)).
		AddRoute(gov.RouterKey, gov.NewHandler(app.govKeeper)).
		AddRoute("ibc", ibc.NewHandler(app.ibcMapper, app.bankKeeper)).
		AddRoute(vpn.RouterKey, vpn.NewHandler(app.vpnKeeper, app.accountKeeper, app.bankKeeper))

	app.QueryRouter().
		AddRoute(staking.QuerierRoute, staking.NewQuerier(app.stakingKeeper, app.cdc)).
		AddRoute(slashing.QuerierRoute, slashing.NewQuerier(app.slashingKeeper, app.cdc)).
		AddRoute(distribution.QuerierRoute, distribution.NewQuerier(app.distributionKeeper)).
		AddRoute(gov.QuerierRoute, gov.NewQuerier(app.govKeeper)).
		AddRoute(vpn.QuerierRoute, vpn.NewQuerier(app.vpnKeeper, app.cdc))

	app.MountStores(app.keyMain, app.keyParams,
		app.keyAccount, app.keyFeeCollection,
		app.keyStaking, app.keySlashing,
		app.keyDistribution, app.keyGov, app.keyMint,
		app.keyIBC, app.keyVPNNode, app.keyVPNSession,
		app.tkeyParams, app.tkeyStaking, app.tkeyDistribution)
	app.SetInitChainer(app.initChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper))
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		if err := app.LoadLatestVersion(app.keyMain); err != nil {
			common.Exit(err.Error())
		}
	}

	return app
}

func MakeCodec() *codec.Codec {
	var cdc = codec.New()
	codec.RegisterCrypto(cdc)
	csdkTypes.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	staking.RegisterCodec(cdc)
	slashing.RegisterCodec(cdc)
	distribution.RegisterCodec(cdc)
	gov.RegisterCodec(cdc)
	ibc.RegisterCodec(cdc)

	vpn.RegisterCodec(cdc)
	return cdc
}

func (app *VPN) BeginBlocker(ctx csdkTypes.Context, req abciTypes.RequestBeginBlock) abciTypes.ResponseBeginBlock {
	mint.BeginBlocker(ctx, app.mintKeeper)
	distribution.BeginBlocker(ctx, req, app.distributionKeeper)
	tags := slashing.BeginBlocker(ctx, req, app.slashingKeeper)

	return abciTypes.ResponseBeginBlock{
		Tags: tags.ToKVPairs(),
	}
}

func (app *VPN) EndBlocker(ctx csdkTypes.Context, req abciTypes.RequestEndBlock) abciTypes.ResponseEndBlock {
	tags := gov.EndBlocker(ctx, app.govKeeper)
	validatorUpdates, endBlockerTags := staking.EndBlocker(ctx, app.stakingKeeper)
	tags = append(tags, endBlockerTags...)
	vpn.EndBlock(ctx, app.vpnKeeper, app.bankKeeper)

	app.assertRuntimeInvariants()

	return abciTypes.ResponseEndBlock{
		ValidatorUpdates: validatorUpdates,
		Tags:             tags,
	}
}

func (app *VPN) initFromGenesisState(ctx csdkTypes.Context, genesisState GenesisState) []abciTypes.ValidatorUpdate {
	genesisState.Sanitize()

	for _, gacc := range genesisState.Accounts {
		acc := gacc.ToAccount()
		acc = app.accountKeeper.NewAccount(ctx, acc)
		app.accountKeeper.SetAccount(ctx, acc)
	}

	distribution.InitGenesis(ctx, app.distributionKeeper, genesisState.DistrData)

	validators, err := staking.InitGenesis(ctx, app.stakingKeeper, genesisState.StakingData)
	if err != nil {
		panic(err)
	}

	auth.InitGenesis(ctx, app.accountKeeper, app.feeCollectionKeeper, genesisState.AuthData)
	bank.InitGenesis(ctx, app.bankKeeper, genesisState.BankData)
	slashing.InitGenesis(ctx, app.slashingKeeper, genesisState.SlashingData, genesisState.StakingData.Validators.ToSDKValidators())
	gov.InitGenesis(ctx, app.govKeeper, genesisState.GovData)
	mint.InitGenesis(ctx, app.mintKeeper, genesisState.MintData)

	if err = VPNValidateGenesisState(genesisState); err != nil {
		panic(err)
	}

	if len(genesisState.GenTxs) > 0 {
		for _, genTx := range genesisState.GenTxs {
			var tx auth.StdTx
			if err = app.cdc.UnmarshalJSON(genTx, &tx); err != nil {
				panic(err)
			}
			bz := app.cdc.MustMarshalBinaryLengthPrefixed(tx)
			res := app.BaseApp.DeliverTx(bz)
			if !res.IsOK() {
				panic(res.Log)
			}
		}

		validators = app.stakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	}

	return validators
}

func (app *VPN) initChainer(ctx csdkTypes.Context, req abciTypes.RequestInitChain) abciTypes.ResponseInitChain {
	stateJSON := req.AppStateBytes

	var genesisState GenesisState
	if err := app.cdc.UnmarshalJSON(stateJSON, &genesisState); err != nil {
		panic(err)
	}

	validators := app.initFromGenesisState(ctx, genesisState)

	if len(req.Validators) > 0 {
		if len(req.Validators) != len(validators) {
			panic(fmt.Errorf("len(RequestInitChain.Validators) != len(validators) (%d != %d)",
				len(req.Validators), len(validators)))
		}
		sort.Sort(abciTypes.ValidatorUpdates(req.Validators))
		sort.Sort(abciTypes.ValidatorUpdates(validators))
		for i, val := range validators {
			if !val.Equal(req.Validators[i]) {
				panic(fmt.Errorf("validators[%d] != req.Validators[%d] ", i, i))
			}
		}
	}

	app.assertRuntimeInvariants()

	return abciTypes.ResponseInitChain{
		Validators: validators,
	}
}

func (app *VPN) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keyMain)
}

type StakingHooks struct {
	dh distribution.Hooks
	sh slashing.Hooks
}

func NewStakingHooks(dh distribution.Hooks, sh slashing.Hooks) StakingHooks {
	return StakingHooks{dh, sh}
}

func (h StakingHooks) AfterValidatorCreated(ctx csdkTypes.Context, valAddr csdkTypes.ValAddress) {
	h.dh.AfterValidatorCreated(ctx, valAddr)
	h.sh.AfterValidatorCreated(ctx, valAddr)
}

func (h StakingHooks) BeforeValidatorModified(ctx csdkTypes.Context, valAddr csdkTypes.ValAddress) {
	h.dh.BeforeValidatorModified(ctx, valAddr)
	h.sh.BeforeValidatorModified(ctx, valAddr)
}

func (h StakingHooks) AfterValidatorRemoved(ctx csdkTypes.Context, consAddr csdkTypes.ConsAddress, valAddr csdkTypes.ValAddress) {
	h.dh.AfterValidatorRemoved(ctx, consAddr, valAddr)
	h.sh.AfterValidatorRemoved(ctx, consAddr, valAddr)
}

func (h StakingHooks) AfterValidatorBonded(ctx csdkTypes.Context, consAddr csdkTypes.ConsAddress, valAddr csdkTypes.ValAddress) {
	h.dh.AfterValidatorBonded(ctx, consAddr, valAddr)
	h.sh.AfterValidatorBonded(ctx, consAddr, valAddr)
}

func (h StakingHooks) AfterValidatorBeginUnbonding(ctx csdkTypes.Context, consAddr csdkTypes.ConsAddress, valAddr csdkTypes.ValAddress) {
	h.dh.AfterValidatorBeginUnbonding(ctx, consAddr, valAddr)
	h.sh.AfterValidatorBeginUnbonding(ctx, consAddr, valAddr)
}

func (h StakingHooks) BeforeDelegationCreated(ctx csdkTypes.Context, delAddr csdkTypes.AccAddress, valAddr csdkTypes.ValAddress) {
	h.dh.BeforeDelegationCreated(ctx, delAddr, valAddr)
	h.sh.BeforeDelegationCreated(ctx, delAddr, valAddr)
}

func (h StakingHooks) BeforeDelegationSharesModified(ctx csdkTypes.Context, delAddr csdkTypes.AccAddress, valAddr csdkTypes.ValAddress) {
	h.dh.BeforeDelegationSharesModified(ctx, delAddr, valAddr)
	h.sh.BeforeDelegationSharesModified(ctx, delAddr, valAddr)
}

func (h StakingHooks) BeforeDelegationRemoved(ctx csdkTypes.Context, delAddr csdkTypes.AccAddress, valAddr csdkTypes.ValAddress) {
	h.dh.BeforeDelegationRemoved(ctx, delAddr, valAddr)
	h.sh.BeforeDelegationRemoved(ctx, delAddr, valAddr)
}

func (h StakingHooks) AfterDelegationModified(ctx csdkTypes.Context, delAddr csdkTypes.AccAddress, valAddr csdkTypes.ValAddress) {
	h.dh.AfterDelegationModified(ctx, delAddr, valAddr)
	h.sh.AfterDelegationModified(ctx, delAddr, valAddr)
}

func (h StakingHooks) BeforeValidatorSlashed(ctx csdkTypes.Context, valAddr csdkTypes.ValAddress, fraction csdkTypes.Dec) {
	h.dh.BeforeValidatorSlashed(ctx, valAddr, fraction)
	h.sh.BeforeValidatorSlashed(ctx, valAddr, fraction)
}
