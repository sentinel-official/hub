package vpn

import (
	"fmt"

	"github.com/sentinel-official/hub/x/deposit"
	"github.com/sentinel-official/hub/x/node"
	"github.com/sentinel-official/hub/x/plan"
	"github.com/sentinel-official/hub/x/provider"
	"github.com/sentinel-official/hub/x/session"
	"github.com/sentinel-official/hub/x/subscription"
	"github.com/sentinel-official/hub/x/vpn/types"
)

func init() {
	node.ParamsSubspace = fmt.Sprintf("%s/%s", types.ModuleName, node.ModuleName)
	subscription.ParamsSubspace = fmt.Sprintf("%s/%s", types.ModuleName, subscription.ModuleName)
	session.ParamsSubspace = fmt.Sprintf("%s/%s", types.ModuleName, session.ModuleName)

	deposit.RouterKey = types.ModuleName
	provider.RouterKey = types.ModuleName
	node.RouterKey = types.ModuleName
	plan.RouterKey = types.ModuleName
	subscription.RouterKey = types.ModuleName
	session.RouterKey = types.ModuleName

	deposit.StoreKey = types.ModuleName
	provider.StoreKey = types.ModuleName
	node.StoreKey = types.ModuleName
	plan.StoreKey = types.ModuleName
	subscription.StoreKey = types.ModuleName
	session.StoreKey = types.ModuleName
}
