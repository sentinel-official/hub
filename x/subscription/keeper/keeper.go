package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	params "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/x/subscription/expected"
	"github.com/sentinel-official/hub/x/subscription/types"
)

type Keeper struct {
	cdc     codec.BinaryMarshaler
	key     sdk.StoreKey
	params  params.Subspace
	bank    expected.BankKeeper
	deposit expected.DepositKeeper
	node    expected.NodeKeeper
	plan    expected.PlanKeeper
}

func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey, params params.Subspace) Keeper {
	return Keeper{
		cdc:    cdc,
		key:    key,
		params: params.WithKeyTable(types.ParamsKeyTable()),
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

func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", types.ModuleName)
}

func (k *Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}
