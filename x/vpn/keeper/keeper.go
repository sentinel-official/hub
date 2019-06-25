package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"

	"github.com/sentinel-official/sentinel-hub/x/deposit"
)

type Keeper struct {
	nodeStoreKey         sdk.StoreKey
	subscriptionStoreKey sdk.StoreKey
	sessionStoreKey      sdk.StoreKey
	cdc                  *codec.Codec
	paramStore           params.Subspace
	depositKeeper        deposit.Keeper
}

func NewKeeper(cdc *codec.Codec, nodeKey, subscriptionStoreKey, sessionKey sdk.StoreKey,
	paramStore params.Subspace, depositKeeper deposit.Keeper) Keeper {

	return Keeper{
		nodeStoreKey:         nodeKey,
		subscriptionStoreKey: subscriptionStoreKey,
		sessionStoreKey:      sessionKey,
		cdc:                  cdc,
		paramStore:           paramStore.WithKeyTable(ParamKeyTable()),
		depositKeeper:        depositKeeper,
	}
}
