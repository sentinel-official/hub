package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"

	"github.com/sentinel-official/hub/x/deposit"
)

type Keeper struct {
	nodeKey         sdk.StoreKey
	subscriptionKey sdk.StoreKey
	resolverKey     sdk.StoreKey
	sessionKey      sdk.StoreKey
	cdc             *codec.Codec
	paramStore      params.Subspace
	deposit         deposit.Keeper
}

func NewKeeper(cdc *codec.Codec, nodeKey, subscriptionKey, sessionKey, resolverKey sdk.StoreKey,
	paramStore params.Subspace, dk deposit.Keeper) Keeper {
	return Keeper{
		nodeKey:         nodeKey,
		subscriptionKey: subscriptionKey,
		sessionKey:      sessionKey,
		resolverKey:     resolverKey,
		cdc:             cdc,
		paramStore:      paramStore.WithKeyTable(ParamKeyTable()),
		deposit:         dk,
	}
}
