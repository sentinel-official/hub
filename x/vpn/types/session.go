package types

import (
	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Session struct {
	ID               sdkTypes.ID        `json:"id"`
	SubscriptionID   sdkTypes.ID        `json:"subscription_id"`
	Bandwidth        sdkTypes.Bandwidth `json:"bandwidth"`
	NodeSign         []byte             `json:"node_sign"`
	ClientSign       []byte             `json:"client_sign"`
	Status           string             `json:"status"`
	StatusModifiedAt int64              `json:"status_modified_at"`
}
