package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/sentinel-official/hub/x/session/expected"
	"github.com/sentinel-official/hub/x/session/types"
)

type Keeper struct {
	cdc          *codec.Codec
	key          sdk.StoreKey
	params       params.Subspace
	plan         expected.PlanKeeper
	subscription expected.SubscriptionKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, params params.Subspace) Keeper {
	return Keeper{
		cdc:    cdc,
		key:    key,
		params: params.WithKeyTable(types.ParamsKeyTable()),
	}
}

func (k *Keeper) WithPlanKeeper(keeper expected.PlanKeeper) {
	k.plan = keeper
}

func (k *Keeper) WithSubscriptionKeeper(keeper expected.SubscriptionKeeper) {
	k.subscription = keeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}
