package types

import (
	"github.com/sentinel-official/hub/x/dvpn/node"
	"github.com/sentinel-official/hub/x/dvpn/provider"
)

type GenesisState struct {
	Providers provider.GenesisState `json:"providers"`
	Nodes     node.GenesisState     `json:"nodes"`
}

func NewGenesisState(providers provider.GenesisState, nodes node.GenesisState) GenesisState {
	return GenesisState{
		Providers: providers,
		Nodes:     nodes,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Providers: provider.DefaultGenesisState(),
		Nodes:     node.DefaultGenesisState(),
	}
}
