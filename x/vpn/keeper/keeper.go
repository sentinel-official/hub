package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	csdkTypes "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	NodeStoreKey    csdkTypes.StoreKey
	SessionStoreKey csdkTypes.StoreKey
	cdc             *codec.Codec
}

func NewKeeper(cdc *codec.Codec, nodeKey, sessionKey csdkTypes.StoreKey) Keeper {
	return Keeper{
		NodeStoreKey:    nodeKey,
		SessionStoreKey: sessionKey,
		cdc:             cdc,
	}
}
