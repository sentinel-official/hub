package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/dvpn/node"
)

type Keeper struct {
	cdc  *codec.Codec
	key  sdk.StoreKey
	node node.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, node node.Keeper) Keeper {
	return Keeper{
		cdc:  cdc,
		key:  key,
		node: node,
	}
}

func (k Keeper) PlanStore(ctx sdk.Context) sdk.KVStore {
	return prefix.NewStore(ctx.KVStore(k.key), []byte("plan/"))
}
