package types

import (
	node "github.com/sentinel-official/hub/x/dvpn/node/types"
	provider "github.com/sentinel-official/hub/x/dvpn/provider/types"
	subscription "github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func init() {
	provider.RouterKey = ModuleName
	node.RouterKey = ModuleName
	subscription.RouterKey = ModuleName

	provider.StoreKey = ModuleName
	node.StoreKey = ModuleName
	subscription.StoreKey = ModuleName
}
