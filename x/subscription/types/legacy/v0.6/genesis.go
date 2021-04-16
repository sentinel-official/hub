package v0_6

import (
	"github.com/sentinel-official/hub/x/subscription/types"
	legacy "github.com/sentinel-official/hub/x/subscription/types/legacy/v0.5"
)

func MigrateGenesisSubscription(item legacy.GenesisSubscription) types.GenesisSubscription {
	return types.GenesisSubscription{
		Subscription: MigrateSubscription(item.Subscription),
		Quotas:       MigrateQuotas(item.Quotas),
	}
}

func MigrateGenesisSubscriptions(items legacy.GenesisSubscriptions) types.GenesisSubscriptions {
	var subscriptions types.GenesisSubscriptions
	for _, item := range items {
		subscriptions = append(subscriptions, MigrateGenesisSubscription(item))
	}

	return subscriptions
}

func MigrateGenesisState(state legacy.GenesisState) types.GenesisState {
	return types.NewGenesisState(
		MigrateGenesisSubscriptions(state.Subscriptions),
		MigrateParams(state.Params),
	)
}
