package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/deposit/expected"
)

type Keeper struct {
	key    sdk.StoreKey
	cdc    *codec.Codec
	supply expected.SupplyKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		key: key,
		cdc: cdc,
	}
}

func (k *Keeper) WithSupplyKeeper(keeper expected.SupplyKeeper) {
	k.supply = keeper
}

func (k Keeper) Store(ctx sdk.Context) sdk.KVStore {
	return prefix.NewStore(ctx.KVStore(k.key), []byte("deposit/"))
}
