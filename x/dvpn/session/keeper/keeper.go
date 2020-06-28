package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/session/expected"
)

type Keeper struct {
	cdc          *codec.Codec
	key          sdk.StoreKey
	subscription expected.SubscriptionKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		cdc: cdc,
		key: key,
	}
}

func (k *Keeper) WithSubscriptionKeeper(keeper expected.SubscriptionKeeper) {
	k.subscription = keeper
}

func (k Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return prefix.NewStore(ctx.KVStore(k.key), []byte("session/"))
}
