package types

import (
	"encoding/hex"
	"fmt"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
)

type Session struct {
	ID               sdkTypes.ID        `json:"id"`
	SubscriptionID   sdkTypes.ID        `json:"subscription_id"`
	Bandwidth        sdkTypes.Bandwidth `json:"bandwidth"`
	NodeOwnerSign    []byte             `json:"node_owner_sign"`
	ClientSign       []byte             `json:"client_sign"`
	Status           string             `json:"status"`
	StatusModifiedAt int64              `json:"status_modified_at"`
}

func (s Session) String() string {
	nodeOwnerSign := hex.EncodeToString(s.NodeOwnerSign)
	clientSign := hex.EncodeToString(s.ClientSign)

	return fmt.Sprintf(`Session
  ID:                   %s
  Subscription ID:      %s
  Bandwidth:            %s
  Node Owner Signature: %s
  Client Signature:     %s
  Status:               %s
  Status Modified At:   %d`, s.ID, s.SubscriptionID, s.Bandwidth,
		nodeOwnerSign, clientSign, s.Status, s.StatusModifiedAt)
}
