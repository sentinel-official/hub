package upgrades

import (
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	ibcicacontrollertypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v4/modules/apps/29-fee/types"
)

const (
	Name = "v11"
)

var (
	StoreUpgrades = &storetypes.StoreUpgrades{
		Added: []string{
			ibcicacontrollertypes.StoreKey,
			ibcfeetypes.StoreKey,
		},
	}
)
