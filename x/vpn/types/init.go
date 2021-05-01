package types

import (
	"fmt"

	deposittypes "github.com/sentinel-official/hub/x/deposit/types"
	nodetypes "github.com/sentinel-official/hub/x/node/types"
	plantypes "github.com/sentinel-official/hub/x/plan/types"
	providertypes "github.com/sentinel-official/hub/x/provider/types"
	sessiontypes "github.com/sentinel-official/hub/x/session/types"
	subscriptiontypes "github.com/sentinel-official/hub/x/subscription/types"
)

func init() {
	nodetypes.ParamsSubspace = fmt.Sprintf("%s/%s", ModuleName, nodetypes.ModuleName)
	subscriptiontypes.ParamsSubspace = fmt.Sprintf("%s/%s", ModuleName, subscriptiontypes.ModuleName)
	sessiontypes.ParamsSubspace = fmt.Sprintf("%s/%s", ModuleName, sessiontypes.ModuleName)

	deposittypes.RouterKey = ModuleName
	providertypes.RouterKey = ModuleName
	nodetypes.RouterKey = ModuleName
	plantypes.RouterKey = ModuleName
	subscriptiontypes.RouterKey = ModuleName
	sessiontypes.RouterKey = ModuleName

	deposittypes.StoreKey = ModuleName
	providertypes.StoreKey = ModuleName
	nodetypes.StoreKey = ModuleName
	plantypes.StoreKey = ModuleName
	subscriptiontypes.StoreKey = ModuleName
	sessiontypes.StoreKey = ModuleName
}
