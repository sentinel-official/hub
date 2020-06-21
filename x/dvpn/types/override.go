package types

import (
	node "github.com/sentinel-official/hub/x/dvpn/node/types"
	provider "github.com/sentinel-official/hub/x/dvpn/provider/types"
)

func init() {
	provider.RouterKey = ModuleName
	node.RouterKey = ModuleName
}
