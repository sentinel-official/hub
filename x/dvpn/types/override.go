package types

import (
	deposit "github.com/sentinel-official/hub/x/dvpn/deposit/types"
	node "github.com/sentinel-official/hub/x/dvpn/node/types"
	provider "github.com/sentinel-official/hub/x/dvpn/provider/types"
	subscription "github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

func init() {
	deposit.RouterKey = ModuleName
	provider.RouterKey = ModuleName
	node.RouterKey = ModuleName
	subscription.RouterKey = ModuleName

	deposit.StoreKey = ModuleName
	provider.StoreKey = ModuleName
	node.StoreKey = ModuleName
	subscription.StoreKey = ModuleName
}
