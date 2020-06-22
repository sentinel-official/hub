package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/provider"
)

type Keeper struct {
	cdc      *codec.Codec
	key      sdk.StoreKey
	provider provider.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, provider provider.Keeper) Keeper {
	return Keeper{
		cdc:      cdc,
		key:      key,
		provider: provider,
	}
}

func (k Keeper) PlanStore(ctx sdk.Context) sdk.KVStore {
	return prefix.NewStore(ctx.KVStore(k.key), []byte("plan/"))
}
