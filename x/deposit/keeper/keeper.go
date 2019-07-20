package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

type Keeper struct {
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	bankKeeper bank.Keeper
}

func NewKeeper(cdc *codec.Codec, storeKey sdk.StoreKey, bankKeeper bank.Keeper) Keeper {
	return Keeper{
		storeKey:   storeKey,
		cdc:        cdc,
		bankKeeper: bankKeeper,
	}
}
