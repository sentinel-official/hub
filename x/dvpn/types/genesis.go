package types

import (
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
)

type GenesisState struct {
	Providers    provider.GenesisState     `json:"providers"`
	Nodes        node.GenesisState         `json:"nodes"`
	Subscription subscription.GenesisState `json:"subscription"`
}

func NewGenesisState(providers provider.GenesisState, nodes node.GenesisState,
	subscription subscription.GenesisState) GenesisState {
	return GenesisState{
		Providers:    providers,
		Nodes:        nodes,
		Subscription: subscription,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Providers:    provider.DefaultGenesisState(),
		Nodes:        node.DefaultGenesisState(),
		Subscription: subscription.DefaultGenesisState(),
	}
}
