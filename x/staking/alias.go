package staking

import (
	"github.com/sentinel-official/hub/x/staking/keeper"
)

// nolint: gochecknoglobals
var (
	RegisterInvariants = keeper.RegisterInvariants
	SupplyInvariants   = keeper.SupplyInvariants
)
