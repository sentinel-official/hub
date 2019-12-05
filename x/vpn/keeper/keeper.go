package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"

	"github.com/sentinel-official/hub/x/deposit"
)

type Keeper struct {
	nodeKey         sdk.StoreKey
	freeClientKey   sdk.StoreKey
	subscriptionKey sdk.StoreKey
	sessionKey      sdk.StoreKey
	cdc             *codec.Codec
	paramStore      params.Subspace
	deposit         deposit.Keeper
}

func NewKeeper(cdc *codec.Codec, nodeKey, freeClientKey, subscriptionKey, sessionKey sdk.StoreKey,
	paramStore params.Subspace, dk deposit.Keeper) Keeper {
	return Keeper{
		nodeKey:         nodeKey,
		freeClientKey:   freeClientKey,
		subscriptionKey: subscriptionKey,
		sessionKey:      sessionKey,
		cdc:             cdc,
		paramStore:      paramStore.WithKeyTable(ParamKeyTable()),
		deposit:         dk,
	}
}
