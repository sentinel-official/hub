package types

import (
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
	"github.com/sentinel-official/hub/x/dvpn/subscription"
)

type GenesisState struct {
	Providers     provider.GenesisState     `json:"providers"`
	Nodes         node.GenesisState         `json:"nodes"`
	Subscriptions subscription.GenesisState `json:"subscriptions"`
}

func NewGenesisState(providers provider.GenesisState, nodes node.GenesisState,
	subscriptions subscription.GenesisState) GenesisState {
	return GenesisState{
		Providers:     providers,
		Nodes:         nodes,
		Subscriptions: subscriptions,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Providers:     provider.DefaultGenesisState(),
		Nodes:         node.DefaultGenesisState(),
		Subscriptions: subscription.DefaultGenesisState(),
	}
}
