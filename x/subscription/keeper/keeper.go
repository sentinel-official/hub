package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/x/subscription/expected"
	"github.com/sentinel-official/hub/x/subscription/types"
)

type Keeper struct {
	appCodec codec.BinaryCodec
	key      sdk.StoreKey
	params   paramstypes.Subspace
	bank     expected.BankKeeper
	deposit  expected.DepositKeeper
	node     expected.NodeKeeper
	plan     expected.PlanKeeper
	session  expected.SessionKeeper
}

func NewKeeper(appCodec codec.BinaryCodec, key sdk.StoreKey, params paramstypes.Subspace) Keeper {
	return Keeper{
		appCodec: appCodec,
		key:      key,
		params:   params.WithKeyTable(types.ParamsKeyTable()),
	}
}

func (k *Keeper) WithBankKeeper(keeper expected.BankKeeper) {
	k.bank = keeper
}

func (k *Keeper) WithDepositKeeper(keeper expected.DepositKeeper) {
	k.deposit = keeper
}

func (k *Keeper) WithNodeKeeper(keeper expected.NodeKeeper) {
	k.node = keeper
}

func (k *Keeper) WithPlanKeeper(keeper expected.PlanKeeper) {
	k.plan = keeper
}

func (k *Keeper) WithSessionKeeper(keeper expected.SessionKeeper) {
	k.session = keeper
}

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}
