package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"

	"github.com/ironman0x7b2/sentinel-sdk/x/deposit"
)

type Keeper struct {
	nodeStoreKey         csdkTypes.StoreKey
	subscriptionStoreKey csdkTypes.StoreKey
	sessionStoreKey      csdkTypes.StoreKey
	cdc                  *codec.Codec
	paramStore           params.Subspace
	depositKeeper        deposit.Keeper
}

func NewKeeper(cdc *codec.Codec, nodeKey, subscriptionStoreKey, sessionKey csdkTypes.StoreKey,
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
