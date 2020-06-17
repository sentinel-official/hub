package types

import (
	provider "github.com/sentinel-official/hub/x/dvpn/provider/types"
)

func init() {
	provider.RouterKey = ModuleName
}
