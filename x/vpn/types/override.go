package types

import (
	deposit "github.com/sentinel-official/hub/x/vpn/deposit/types"
	node "github.com/sentinel-official/hub/x/vpn/node/types"
	plan "github.com/sentinel-official/hub/x/vpn/plan/types"
	provider "github.com/sentinel-official/hub/x/vpn/provider/types"
	session "github.com/sentinel-official/hub/x/vpn/session/types"
	subscription "github.com/sentinel-official/hub/x/vpn/subscription/types"
)

func init() {
	deposit.RouterKey = ModuleName
	provider.RouterKey = ModuleName
	node.RouterKey = ModuleName
	plan.RouterKey = ModuleName
	subscription.RouterKey = ModuleName
	session.RouterKey = ModuleName

	deposit.StoreKey = ModuleName
	provider.StoreKey = ModuleName
	node.StoreKey = ModuleName
	plan.StoreKey = ModuleName
	subscription.StoreKey = ModuleName
	session.StoreKey = ModuleName
}
