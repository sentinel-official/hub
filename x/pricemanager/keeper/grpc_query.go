package keeper

import (
	"github.com/sentinel-official/hub/v12/x/pricemanager/types"
)

var _ types.QueryServer = Keeper{}
