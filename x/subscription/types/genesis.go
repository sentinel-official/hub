package types

import (
	"fmt"
)

type (
	GenesisSubscriptions []GenesisSubscription
)

func NewGenesisState(subscriptions GenesisSubscriptions, params Params) *GenesisState {
	return &GenesisState{
		Subscriptions: subscriptions,
		Params:        params,
	}
}

func DefaultGenesisState() *GenesisState {
	return NewGenesisState(nil, DefaultParams())
}

func ValidateGenesis(state *GenesisState) error {
	if err := state.Params.Validate(); err != nil {
		return err
	}

	subscriptions := make(map[uint64]bool)
	for _, item := range state.Subscriptions {
		if subscriptions[item.Subscription.Id] {
			return fmt.Errorf("found duplicate subscription for id %d", item.Subscription.Id)
		}

		subscriptions[item.Subscription.Id] = true
	}

	for _, item := range state.Subscriptions {
		quotas := make(map[string]bool)
		for _, quota := range item.Quotas {
			if quotas[quota.Address] {
				return fmt.Errorf("found duplicate quota for subscription %d and address %s", item.Subscription.Id, quota.Address)
			}

			quotas[quota.Address] = true
		}
	}

	for _, item := range state.Subscriptions {
		if err := item.Subscription.Validate(); err != nil {
			return err
		}
	}

	for _, item := range state.Subscriptions {
		for _, quota := range item.Quotas {
			if err := quota.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}
