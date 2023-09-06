// DO NOT COVER

package types

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"

	deposittypes "github.com/sentinel-official/hub/v1/x/deposit/types"
	nodetypes "github.com/sentinel-official/hub/v1/x/node/types"
	plantypes "github.com/sentinel-official/hub/v1/x/plan/types"
	providertypes "github.com/sentinel-official/hub/v1/x/provider/types"
	sessiontypes "github.com/sentinel-official/hub/v1/x/session/types"
	subscriptiontypes "github.com/sentinel-official/hub/v1/x/subscription/types"
)

func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	deposittypes.RegisterInterfaces(registry)
	providertypes.RegisterInterfaces(registry)
	nodetypes.RegisterInterfaces(registry)
	plantypes.RegisterInterfaces(registry)
	subscriptiontypes.RegisterInterfaces(registry)
	sessiontypes.RegisterInterfaces(registry)
}
