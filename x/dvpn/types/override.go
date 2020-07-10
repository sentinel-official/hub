package types

import (
	deposit "github.com/sentinel-official/hub/x/dvpn/deposit/types"
	node "github.com/sentinel-official/hub/x/dvpn/node/types"
	plan "github.com/sentinel-official/hub/x/dvpn/plan/types"
	provider "github.com/sentinel-official/hub/x/dvpn/provider/types"
	session "github.com/sentinel-official/hub/x/dvpn/session/types"
	subscription "github.com/sentinel-official/hub/x/dvpn/subscription/types"
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
