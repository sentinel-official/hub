package v0_5

import (
	deposit "github.com/sentinel-official/hub/x/deposit/types/legacy/v0.5"
	node "github.com/sentinel-official/hub/x/node/types/legacy/v0.5"
	plan "github.com/sentinel-official/hub/x/plan/types/legacy/v0.5"
	provider "github.com/sentinel-official/hub/x/provider/types/legacy/v0.5"
	session "github.com/sentinel-official/hub/x/session/types/legacy/v0.5"
	subscription "github.com/sentinel-official/hub/x/subscription/types/legacy/v0.5"
)

type (
	GenesisState struct {
		Deposits      deposit.GenesisState       `json:"deposits"`
		Providers     *provider.GenesisState     `json:"providers"`
		Nodes         *node.GenesisState         `json:"nodes"`
		Plans         plan.GenesisState          `json:"plans"`
		Subscriptions *subscription.GenesisState `json:"subscriptions"`
		Sessions      *session.GenesisState      `json:"sessions"`
	}
)
