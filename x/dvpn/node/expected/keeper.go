package expected

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	hub "github.com/sentinel-official/hub/types"
	provider "github.com/sentinel-official/hub/x/dvpn/provider/types"
	subscription "github.com/sentinel-official/hub/x/dvpn/subscription/types"
)

type ProviderKeeper interface {
	GetProvider(ctx sdk.Context, address hub.ProvAddress) (provider.Provider, bool)
}

type SubscriptionKeeper interface {
	GetPlansForProvider(ctx sdk.Context, address hub.ProvAddress) subscription.Plans
	DeleteNodeAddressForPlan(ctx sdk.Context, id uint64, address hub.NodeAddress)
}
