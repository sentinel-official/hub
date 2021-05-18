package v05

type (
	GenesisSubscription struct {
		Subscription Subscription `json:"_"`
		Quotas       Quotas       `json:"quotas"`
	}

	GenesisSubscriptions []GenesisSubscription

	GenesisState struct {
		Subscriptions GenesisSubscriptions `json:"_"`
		Params        Params               `json:"params"`
	}
)
