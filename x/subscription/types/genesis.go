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

	m := make(map[uint64]bool)
	for _, item := range state.Subscriptions {
		id := item.Subscription.GetCachedValue().(*BaseSubscription).ID
		if m[id] {
			return fmt.Errorf("found a duplicate subscription for id %d", id)
		}

		m[id] = true
	}

	for _, item := range state.Subscriptions {
		var (
			m  = make(map[string]bool)
			id = item.Subscription.GetCachedValue().(*BaseSubscription).ID
		)

		for _, quota := range item.Quotas {
			if m[quota.Address] {
				return fmt.Errorf("found a duplicate quota %d/%s", id, quota.Address)
			}

			m[quota.Address] = true
		}
	}

	for _, item := range state.Subscriptions {
		item := item.Subscription.GetCachedValue().(Subscription)
		if err := item.Validate(); err != nil {
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
