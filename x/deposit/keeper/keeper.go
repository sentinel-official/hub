package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

type Keeper struct {
	storeKey   csdkTypes.StoreKey
	cdc        *codec.Codec
	bankKeeper bank.Keeper
}

func NewKeeper(cdc *codec.Codec, storeKey csdkTypes.StoreKey, bankKeeper bank.Keeper) Keeper {
	return Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		bankKeeper: bankKeeper,
	}
}
