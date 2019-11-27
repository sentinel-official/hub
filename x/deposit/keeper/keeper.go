package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/supply"
)

type Keeper struct {
	key    sdk.StoreKey
	cdc    *codec.Codec
	supply supply.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, sk supply.Keeper) Keeper {
	return Keeper{
		key:    key,
		cdc:    cdc,
		supply: sk,
	}
}
