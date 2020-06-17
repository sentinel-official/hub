package types

import (
	"github.com/sentinel-official/hub/x/dvpn/provider"
)

type GenesisState struct {
	Providers provider.GenesisState `json:"providers"`
}

func NewGenesisState(provider provider.GenesisState) GenesisState {
	return GenesisState{
		Providers: provider,
	}
}

func DefaultGenesisState() GenesisState {
	return GenesisState{
		Providers: provider.DefaultGenesisState(),
	}
}
