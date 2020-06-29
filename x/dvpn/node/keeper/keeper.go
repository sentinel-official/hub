package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/node/expected"
	"github.com/sentinel-official/hub/x/dvpn/node/types"
)

type Keeper struct {
	cdc      *codec.Codec
	key      sdk.StoreKey
	provider expected.ProviderKeeper
	plan     expected.PlanKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		cdc: cdc,
		key: key,
	}
}

func (k *Keeper) WithProviderKeeper(keeper expected.ProviderKeeper) {
	k.provider = keeper
}

func (k *Keeper) WithPlanKeeper(keeper expected.PlanKeeper) {
	k.plan = keeper
}

func (k Keeper) Store(ctx sdk.Context) sdk.KVStore {
	child := fmt.Sprintf("%s/", types.ModuleName)
	return prefix.NewStore(ctx.KVStore(k.key), []byte(child))
}
